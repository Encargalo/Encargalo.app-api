package repo

import (
	models "CaliYa/core/domain/models/customers"
	"CaliYa/core/domain/ports"
	"CaliYa/core/errors"
	"context"
	"database/sql"
	"fmt"

	"github.com/uptrace/bun"
)

type customersRepo struct {
	db *bun.DB
}

func NewCustomersRepo(db *bun.DB) ports.CustomersRepo {
	return &customersRepo{db}
}

func (c *customersRepo) RegisterCustomer(ctx context.Context, customer *models.Accounts) error {

	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("error iniciando transacción: %w", err)
	}

	if _, err := tx.NewInsert().Model(customer).Returning("id").Exec(ctx); err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("error al insertar el customer")
	}

	activationAccount := new(models.ActivateAccount)
	activationAccount.BuildActivateAccount(customer.ID)

	if _, err := tx.NewInsert().Model(activationAccount).Exec(ctx); err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("error al registrar el codigo de activación")
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error al confirmar transacción: %w", err)
	}

	return nil
}

func (c *customersRepo) SearchCustomerByPhone(ctx context.Context, phone string) (*models.Accounts, error) {

	account := new(models.Accounts)

	err := c.db.NewSelect().
		Model(account).
		Where("phone = ?", phone).
		Where("deleted_at IS NULL").
		Relation("Addresses").
		Scan(ctx)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrNotFound
		}
		return nil, errors.ErrUnexpected
	}

	return account, nil

}
