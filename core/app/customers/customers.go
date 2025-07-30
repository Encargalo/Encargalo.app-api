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

	customer, err := c.repo.SearchCustomerBy(ctx, criteria)
	if err != nil {
		return nil, err
	}

	return customer, nil

}
