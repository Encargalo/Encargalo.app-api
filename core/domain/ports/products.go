package ports

import (
	"CaliYa/core/domain/models"
	"context"
)

type ProductsApp interface {
	RegisterProduct(ctx context.Context) error

	GetProductByCategory(ctx context.Context, category string) ([]models.Items, error)
}

type ProductsRepo interface {
	GetProductByCategory(ctx context.Context, category string) ([]models.Items, error)
}
