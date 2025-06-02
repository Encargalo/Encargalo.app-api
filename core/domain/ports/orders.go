package ports

import (
	"CaliYa/core/domain/models"
	"context"
)

type OrdersApp interface {
	RegisterOrders(ctx context.Context, order models.Order) error
}

type OrdersRepo interface {
	RegisterOrders(ctx context.Context, order *models.Order) error
	CalculatePrice(ctx context.Context, order *models.Order)
}
