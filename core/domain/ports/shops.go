package ports

import (
	"CaliYa/core/domain/dto"
	"CaliYa/core/domain/models/items"
	"context"
)

type ShopsApp interface {
	GetAllShops(ctx context.Context) (dto.ShopsResponse, error)
	GetShopsBy(ctx context.Context, criteria dto.SearchShopsByID) (*items.ItemsShops, error)
}

type ShopsRepo interface {
	GetAllShops(ctx context.Context) (dto.ShopsResponse, error)
	GetShopsBy(ctx context.Context, criteria dto.SearchShopsByID) (*items.ItemsShops, error)
}
