package repo

import (
	dto "CaliYa/core/domain/dto/customers"
	models "CaliYa/core/domain/models/customers"
	"CaliYa/core/domain/ports"
	"CaliYa/core/errors"
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
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
		fmt.Println("error iniciando transacción: %w", err)
		return errors.ErrUnexpected
	}

	if _, err := tx.NewInsert().Model(customer).Returning("id").Exec(ctx); err != nil {
		_ = tx.Rollback()
		fmt.Println("error al insertar el customer")
		return errors.ErrUnexpected
	}

	activationAccount := new(models.ActivateAccount)
	activationAccount.BuildActivateAccount(customer.ID)

	if _, err := tx.NewInsert().Model(activationAccount).Exec(ctx); err != nil {
		_ = tx.Rollback()
		fmt.Println("error al registrar el codigo de activación")
		return errors.ErrUnexpected
	}

	if err := tx.Commit(); err != nil {
		fmt.Println("error al confirmar transacción: %w", err)
		return errors.ErrUnexpected
	}

	return nil
}

func (c *customersRepo) SearchCustomerBy(ctx context.Context, criteria dto.SearchCustomerBy) (*models.Accounts, error) {

	account := new(models.Accounts)

	err := c.db.NewSelect().
		Model(account).
		WhereGroup("or", func(sq *bun.SelectQuery) *bun.SelectQuery {
			if criteria.ID != uuid.Nil {
				sq = sq.Where("id = ?", criteria.ID)
			}
			if criteria.Phone != "" {
				sq = sq.Where("phone = ?", criteria.Phone)
			}
			return sq
		}).
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
