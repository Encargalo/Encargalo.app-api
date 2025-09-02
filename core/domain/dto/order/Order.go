package order

import (
	"CaliYa/core/domain/dto/customers"
	"context"

	"github.com/go-playground/mold/modifiers"
	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
	conform  = modifiers.New()
)

type CreateOrder struct {
	ShopID         string           `json:"shop_id" validate:"required,uuid4" mold:"trim" example:"0936ea77-72b8-46eb-b80c-c8c22386a0fb"`
	CustomerID     string           `json:"customer_id" validate:"required,uuid4" mold:"trim" example:"84ad5b8c-7877-4a60-b0a1-c6545f6a1344"`
	Address        string           `json:"address" validate:"required,min=10,max=100" example:"123 Main St, Springfield"`
	Coords         customers.Coords `json:"coords" validate:"required"`
	Method_Payment string           `json:"method_payment" validate:"required,oneof=Nequi Efectivo" example:"Nequi"`
	Delivery_Price int              `json:"delivery_price" validate:"gte=0" example:"5000"`
	Total_Price    int              `json:"total_price" swaggerignore:"true"`

	Items []OrderItems `json:"items" validate:"required,min=1"`
}

func (io *CreateOrder) Validate() error {
	_ = conform.Struct(context.Background(), io)
	return validate.Struct(io)
}
