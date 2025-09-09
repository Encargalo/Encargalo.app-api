package repo

import (
	"CaliYa/core/domain/models"
	ordersModels "CaliYa/core/domain/models/orders"
	"CaliYa/core/domain/ports"
	"context"
	"fmt"

	"github.com/uptrace/bun"
)

type orders struct {
	db *bun.DB
}

func NewOrdersRepo(db *bun.DB) ports.OrdersRepo {
	return &orders{db}
}

func (o *orders) RegisterOrders(ctx context.Context, order *ordersModels.Order) error {

	tx, err := o.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("error iniciando transacción: %w", err)
	}

	if _, err := tx.NewInsert().Model(order).Exec(ctx); err != nil {
		return fmt.Errorf("error insertando orden: %w", err)
	}

	if _, err := tx.NewInsert().Model(&order.ItemsOrder).Exec(ctx); err != nil {
		return fmt.Errorf("error insertando items: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error al confirmar transacción: %w", err)
	}

	return nil
}

func (o *orders) CalculatePrice(ctx context.Context, orders *ordersModels.Order) {

	var items []models.Items

	err := o.db.NewSelect().Table("products.items").Where("id in (?)", bun.In(orders.GetItemsID())).Scan(ctx, &items)

	if err != nil {
		fmt.Println(err)
	}

	orders.SetPrices(items)

}
