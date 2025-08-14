package customers

import (
	"CaliYa/core/domain/ports/customers"
	"context"

	customersdto "CaliYa/core/domain/dto/customers"

	customerModel "CaliYa/core/domain/models/customers"

	"github.com/google/uuid"
)

type customersAddressApp struct {
	repo customers.CustomersAddressRepo
}

func NewCustomersAddressApp(repo customers.CustomersAddressRepo) customers.CustomersAddressApp {
	return &customersAddressApp{repo}
}

func (c *customersAddressApp) RegisterAddress(ctx context.Context, customerID uuid.UUID, address customersdto.Address) error {

	addresModel := customerModel.Address{}
	addresModel.BuildToModel(customerID, address)

	return c.repo.RegisterAddress(ctx, addresModel)
}

func (c *customersAddressApp) SearchAllAddress(ctx context.Context, customer_id uuid.UUID) (customersdto.Addresses, error) {

	return c.repo.SearchAllAddress(ctx, customer_id)
}

func (c *customersAddressApp) DeleteAddress(ctx context.Context, customer_id, address_id uuid.UUID) error {
	return c.repo.DeleteAddress(ctx, customer_id, address_id)
}
