package ports

import (
	"CaliYa/core/domain/models"
	"context"
)

type PromotionsApp interface {
	GetPromotions(ctx context.Context) ([]models.ItemsPromo, error)
}

type PromotionsRepository interface {
	GetPromotions(ctx context.Context) ([]models.ItemsPromo, error)
}
