package ports

import (
	"CaliYa/core/domain/dto/order"
	"CaliYa/core/domain/models/orders"
	"context"
)

type OrdersApp interface {
	RegisterOrders(ctx context.Context, order order.CreateOrder) error
}

type OrdersRepo interface {
	RegisterOrders(ctx context.Context, order *orders.Order) error
	CalculatePrice(ctx context.Context, order *orders.Order)
}
