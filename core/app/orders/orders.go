package orders

import (
	"CaliYa/core/domain/dto/order"
	ordersModel "CaliYa/core/domain/models/orders"
	"CaliYa/core/domain/ports"
	"context"
	"fmt"
)

type orders struct {
	repo ports.OrdersRepo
}

func NewOrdersApp(repo ports.OrdersRepo) ports.OrdersApp {
	return &orders{repo}
}

func (o *orders) RegisterOrders(ctx context.Context, order order.CreateOrder) error {

	orderModel := ordersModel.Order{}
	orderModel.BuildDtoToModel(order)

	o.repo.CalculatePrice(ctx, &orderModel)

	if err := o.repo.RegisterOrders(ctx, &orderModel); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
