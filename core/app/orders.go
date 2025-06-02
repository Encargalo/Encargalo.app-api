package app

import (
	"CaliYa/core/domain/models"
	"CaliYa/core/domain/ports"
	"context"
	"fmt"

	"github.com/google/uuid"
)

type orders struct {
	repo ports.OrdersRepo
}

func NewOrdersApp(repo ports.OrdersRepo) ports.OrdersApp {
	return &orders{repo}
}

func (o *orders) RegisterOrders(ctx context.Context, order models.Order) error {

	oID, _ := uuid.NewUUID()

	order.ID = oID

	for i := range order.ItemsOrders {
		order.ItemsOrders[i].OrderID = oID
	}

	if err := o.repo.RegisterOrders(ctx, &order); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
