package ports

import (
	"CaliYa/core/domain/dto"
	"CaliYa/core/domain/models"
	"context"
)

type ProductsApp interface {
	RegisterProduct(ctx context.Context) error
	GetProductsBy(ctx context.Context, criteria dto.SearchProductsBy) (*models.ProductsShops, error)
	GetProductByCategory(ctx context.Context, category string) ([]models.Items, error)
}

type ProductsRepo interface {
	GetProductsBy(ctx context.Context, criteria dto.SearchProductsBy) (*models.ProductsShops, error)
	GetProductByCategory(ctx context.Context, category string) ([]models.Items, error)
}
