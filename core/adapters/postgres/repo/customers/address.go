package customers

import (
	"CaliYa/core/domain/ports/customers"
	"context"
	"database/sql"
	"errors"
	"fmt"

	customersDTO "CaliYa/core/domain/dto/customers"
	customersModel "CaliYa/core/domain/models/customers"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type customersAddressRepo struct {
	db *bun.DB
}

func NewCustomersAddressRepo(db *bun.DB) customers.CustomersAddressRepo {
	return &customersAddressRepo{db}
}

func (c *customersAddressRepo) RegisterAddress(ctx context.Context, address customersModel.Address) error {

	if _, err := c.db.NewInsert().Model(&address).Exec(ctx); err != nil {
		fmt.Println("error al insertar la direccion del customer - ", err.Error())
		return fmt.Errorf("unexpected error")
	}

	return nil
}

func (c *customersAddressRepo) SearchAllAddress(ctx context.Context, customer_id uuid.UUID) (customersDTO.Addresses, error) {

	var addresses customersModel.Addresses

	if err := c.db.NewSelect().Model(&addresses).
		Where("customer_id = ?", customer_id).Scan(ctx); err != nil {
		if err == sql.ErrNoRows {
			return customersDTO.Addresses{}, errors.New("not found")
		}
		fmt.Println(err.Error())
		return customersDTO.Addresses{}, errors.New("unexpected error")

	}

	return addresses.ToDomainDTO(), nil
}
