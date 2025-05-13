package ports

import (
	"CaliYa/core/domain/models"
	"context"
)

type ProductsApp interface {
	RegisterProduct(ctx context.Context) error
	GetProducts(ctx context.Context) (*models.ProductsShops, error)
}

type ProductsRepo interface {
	GetProducts(ctx context.Context) (*models.ProductsShops, error)
}
