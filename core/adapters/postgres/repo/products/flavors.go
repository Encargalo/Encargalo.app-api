package products

import (
	"CaliYa/core/domain/models/items"
	"CaliYa/core/domain/ports/products"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type flavors struct {
	db *bun.DB
}

func NewFlavorRepo(db *bun.DB) products.FlavorsRepo {
	return &flavors{db: db}
}

func (f *flavors) SearchFlavorsByProductID(ctx context.Context, productID uuid.UUID) ([]items.Flavor, error) {
	var flavors []items.Flavor

	if err := f.db.NewSelect().
		Model(&flavors).
		Where("product_id = ?", productID).
		Scan(ctx); err != nil {
		fmt.Println("error querying flavors by product ID: %w", err)
		return nil, errors.New("unexpected error")
	}

	if len(flavors) == 0 {
		return nil, errors.New("not found")
	}

	return flavors, nil
}
