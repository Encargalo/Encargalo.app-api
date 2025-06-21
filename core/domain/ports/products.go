package ports

import (
	"CaliYa/core/domain/models"
	"context"

	"github.com/google/uuid"
)

type ProductsApp interface {
	RegisterProduct(ctx context.Context) error
	GetProductByCategory(ctx context.Context, category string) ([]models.Items, error)
	GetAditionsByCategory(ctx context.Context, id uuid.UUID) ([]models.Items, error)
}

type ProductsRepo interface {
	GetProductByCategory(ctx context.Context, category string) ([]models.Items, error)
	GetAditionsByCategory(ctx context.Context, id uuid.UUID) ([]models.Items, error)
}
