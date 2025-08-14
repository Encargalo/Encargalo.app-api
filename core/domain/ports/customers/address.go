package customers

import (
	"CaliYa/core/domain/dto/customers"
	customersDTO "CaliYa/core/domain/dto/customers"
	customersModel "CaliYa/core/domain/models/customers"
	"context"

	"github.com/google/uuid"
)

type CustomersAddressApp interface {
	RegisterAddress(ctx context.Context, customerID uuid.UUID, address customers.Address) error
	SearchAllAddress(ctx context.Context, customer_id uuid.UUID) (customers.Addresses, error)
	DeleteAddress(ctx context.Context, customer_id, address_id uuid.UUID) error
}

type CustomersAddressRepo interface {
	RegisterAddress(ctx context.Context, address customersModel.Address) error
	SearchAllAddress(ctx context.Context, customer_id uuid.UUID) (customersDTO.Addresses, error)
	DeleteAddress(ctx context.Context, customer_id, address_id uuid.UUID) error
}
