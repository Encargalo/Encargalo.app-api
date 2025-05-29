package ports

import (
	"CaliYa/core/domain/models"
	"context"
)

type ProductsApp interface {
	RegisterProduct(ctx context.Context) error
	GetProductsBy(ctx context.Context, criteria models.SearchProductsBy) (*models.ProductsShops, error)
}

type ProductsRepo interface {
	GetProductsBy(ctx context.Context, criteria models.SearchProductsBy) (*models.ProductsShops, error)
}
