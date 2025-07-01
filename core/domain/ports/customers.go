package ports

import (
	"CaliYa/core/domain/dto"
	models "CaliYa/core/domain/models/customers"
	"context"
)

type CustomersApp interface {
	RegisterCustomer(ctx context.Context, customer dto.RegisterCustomer) error
}

type CustomersRepo interface {
	RegisterCustomer(ctx context.Context, customer *models.Accounts) error
	SearchCustomerByPhone(ctx context.Context, phone string) (*models.Accounts, error)
}
