package app

import (
	dto "CaliYa/core/domain/dto/customers"
	models "CaliYa/core/domain/models/customers"
	"CaliYa/core/domain/ports"
	"CaliYa/core/errors"
	"CaliYa/core/utils"
	"context"

	"fmt"
)

type customersApp struct {
	repo ports.CustomersRepo
	pass utils.Password
}

func NewCustomerApp(repo ports.CustomersRepo, pass utils.Password) ports.CustomersApp {
	return &customersApp{
		repo,
		pass,
	}
}

func (c *customersApp) RegisterCustomer(ctx context.Context, customer dto.RegisterCustomer) error {

	custo, err := c.SearchCustomerBy(ctx, dto.SearchCustomerBy{Phone: customer.Phone})
	if err != nil {
		if err == errors.ErrUnexpected {
			return err
		}
	}

	if custo != nil {
		return fmt.Errorf("phone al ready exist")
	}

	c.pass.HashPassword(&customer.Password)

	customerModel := models.Accounts{}
	customerModel.BuildCustomerRegisterModel(customer)

	_, err = c.repo.RegisterCustomer(ctx, &customerModel)
	if err != nil {
		return err
	}

	return nil
}

func (c *customersApp) SearchCustomerBy(ctx context.Context, criteria dto.SearchCustomerBy) (*models.Accounts, error) {

	customer, err := c.repo.SearchCustomerBy(ctx, criteria)
	if err != nil {
		return nil, err
	}

	return customer, nil

}
