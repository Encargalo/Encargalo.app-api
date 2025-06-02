package ports

import (
	"CaliYa/core/domain/dto"
	"context"
)

type ShopsApp interface {
	GetAllShops(ctx context.Context) (dto.ShopsResponse, error)
}

type ShopsRepo interface {
	GetAllShops(ctx context.Context) (dto.ShopsResponse, error)
}
