package ports

import (
	ItemsModels "CaliYa/core/domain/models/items"
	"context"

	"github.com/google/uuid"
)

type ProductsApp interface {
	RegisterProduct(ctx context.Context) error
	GetProductByCategory(ctx context.Context, category string) ([]ItemsModels.Items, error)
	GetAditionsByCategory(ctx context.Context, id uuid.UUID) ([]ItemsModels.Items, error)
}

type ProductsRepo interface {
	GetProductByCategory(ctx context.Context, category string) ([]ItemsModels.Items, error)
	GetAditionsByCategory(ctx context.Context, id uuid.UUID) ([]ItemsModels.Items, error)
}
