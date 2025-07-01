package app

import (
	"CaliYa/core/domain/dto"
	models "CaliYa/core/domain/models/customers"
	"CaliYa/core/domain/ports"
	"CaliYa/core/utils"
	"context"
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

	c.pass.HashPassword(&customer.Password)

	customerModel := models.Accounts{}
	customerModel.BuildCustomerRegisterModel(customer)

	if err := c.repo.RegisterCustomer(ctx, &customerModel); err != nil {
		return err
	}

	return nil
}
