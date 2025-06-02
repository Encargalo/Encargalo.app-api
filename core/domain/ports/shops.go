package ports

import (
	"CaliYa/core/domain/dto"
	"CaliYa/core/domain/models"
	"context"
)

type ShopsApp interface {
	GetAllShops(ctx context.Context) (dto.ShopsResponse, error)
	GetShopsBy(ctx context.Context, criteria dto.SearchShopsByID) (*models.ProductsShops, error)
}

type ShopsRepo interface {
	GetAllShops(ctx context.Context) (dto.ShopsResponse, error)
	GetShopsBy(ctx context.Context, criteria dto.SearchShopsByID) (*models.ProductsShops, error)
}
