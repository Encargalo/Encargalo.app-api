package models

import (
	"context"

	"github.com/go-playground/mold/modifiers"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

var (
	validate = validator.New()
	conform  = modifiers.New()
)

type Order struct {
	bun.BaseModel `bun:"table:business.orders" swaggerignore:"true"`

	ID          uuid.UUID `bun:"id,pk,type:uuid" swaggerignore:"true"`
	ShopID      uuid.UUID `bun:"shop_id" json:"shop_id" validate:"required,uuid4" mold:"trim" example:"0936ea77-72b8-46eb-b80c-c8c22386a0fb"`
	CustomerID  uuid.UUID `bun:"customer_id" json:"customer_id" validate:"required,uuid4" mold:"trim" example:"84ad5b8c-7877-4a60-b0a1-c6545f6a1344"`
	Total_Price int       `bun:"total_price" swaggerignore:"true"`

	ItemsOrders []ItemsOrders `bun:"rel:has-many,join:id=order_id" json:"items_order" validate:"required"`
}

func (o *Order) Validate() error {
	_ = conform.Struct(context.Background(), o)
	return validate.Struct(o)
}

func (o *Order) GetItemsID() []uuid.UUID {

	ret := make([]uuid.UUID, len(o.ItemsOrders))

	for i := range ret {
		ret[i] = o.ItemsOrders[i].ItemID
	}

	return ret

}

func (o *Order) SetPrices(items []Items) {

	for _, item := range items {
		for i, io := range o.ItemsOrders {
			if io.ItemID == item.ID {

				var subTotal int

				subTotal += item.Price * io.Amount

				o.ItemsOrders[i].TotalPrice = subTotal
				o.ItemsOrders[i].UnitPrice = item.Price

				o.Total_Price += subTotal

			}
		}
	}

}
