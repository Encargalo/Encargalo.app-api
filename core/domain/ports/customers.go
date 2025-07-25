package ports

import (
	dto "CaliYa/core/domain/dto/customers"
	models "CaliYa/core/domain/models/customers"
	"context"

	"github.com/google/uuid"
)

type CustomersApp interface {
	RegisterCustomer(ctx context.Context, customer dto.RegisterCustomer) (uuid.UUID, error)
	SearchCustomerBy(ctx context.Context, criteria dto.SearchCustomerBy) (*models.Accounts, error)
}

type CustomersRepo interface {
	RegisterCustomer(ctx context.Context, customer *models.Accounts) (*models.Accounts, error)
	SearchCustomerBy(ctx context.Context, criteria dto.SearchCustomerBy) (*models.Accounts, error)
}
