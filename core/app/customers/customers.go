package customers

import (
	dto "CaliYa/core/domain/dto/customers"
	models "CaliYa/core/domain/models/customers"
	ports "CaliYa/core/domain/ports/customers"
	sessions "CaliYa/core/domain/ports/sessions"
	calierrors "CaliYa/core/errors"
	"CaliYa/core/utils"
	"context"
	"errors"

	"github.com/google/uuid"
)

type customersApp struct {
	repo        ports.CustomersRepo
	pass        utils.Password
	sessionsSvc sessions.SessionsApp
}

func NewCustomerApp(repo ports.CustomersRepo, pass utils.Password, sessionsSvc sessions.SessionsApp) ports.CustomersApp {
	return &customersApp{
		repo,
		pass,
		sessionsSvc,
	}
}

func (c *customersApp) RegisterCustomer(ctx context.Context, customer dto.RegisterCustomer) (uuid.UUID, error) {

	custo, err := c.SearchCustomerBy(ctx, dto.SearchCustomerBy{Phone: customer.Phone})
	if err != nil {
		if err == calierrors.ErrUnexpected {
			return uuid.Nil, err
		}
	}

	if custo != nil {
		return uuid.Nil, errors.New("phone al ready exist")
	}

	c.pass.HashPassword(&customer.Password)

	customerModel := models.Accounts{}
	customerModel.BuildCustomerRegisterModel(customer)

	custo, err = c.repo.RegisterCustomer(ctx, &customerModel)
	if err != nil {
		return uuid.Nil, err
	}

	sessionID, err := c.sessionsSvc.RegisterSessions(ctx, custo.ID, "Client")
	if err != nil {
		return uuid.Nil, err
	}

	return sessionID, nil
}

func (c *customersApp) SearchCustomerBy(ctx context.Context, criteria dto.SearchCustomerBy) (*models.Accounts, error) {
	return c.repo.SearchCustomerBy(ctx, criteria)
}

func (c *customersApp) UpdateCustomer(ctx context.Context, customer_id uuid.UUID, customer dto.UpdateCustomer) error {

	criteria := dto.SearchCustomerBy{
		ID: customer_id,
	}

	_, err := c.SearchCustomerBy(ctx, criteria)
	if err != nil {
		return err
	}

	cust, err := c.repo.SearchCustomerByPhoneAndNotIDEquals(ctx, customer_id, customer.Phone)
	if err != nil {
		if err.Error() != "not found." {
			return err
		}
	}

	if cust != nil {
		return errors.New("phone al ready exist")
	}

	customerModel := models.Accounts{}
	customerModel.BuildCustomerUpdateModel(customer)

	return c.repo.UpdateCustomer(ctx, customer_id, &customerModel)
}

func (c *customersApp) UpdatePassword(ctx context.Context, customer_id uuid.UUID, pass dto.UpdatePassword) error {

	criteria := dto.SearchCustomerBy{
		ID: customer_id,
	}

	customer, err := c.SearchCustomerBy(ctx, criteria)
	if err != nil {
		return err
	}

	if customer == nil {
		return calierrors.ErrNotFound
	}

	customerModel := models.Accounts{}
	customerModel.BuildCustomerUpdatePasswordModel(pass)

	c.pass.HashPassword(&customerModel.Password)

	return c.repo.UpdatePassword(ctx, customer_id, &customerModel)

}
