package order

import (
	"CaliYa/core/domain/dto/customers"
	"context"

	"github.com/go-playground/mold/modifiers"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var (
	validate = validator.New()
	conform  = modifiers.New()
)

type CreateOrder struct {
	ID             uuid.UUID        `json:"id" validate:"required,uuid4" example:"0936ea77-72b8-46eb-b80c-c8c22386a0fb"`
	ShopID         uuid.UUID        `json:"shop_id" validate:"required,uuid4" mold:"trim" example:"0936ea77-72b8-46eb-b80c-c8c22386a0fb"`
	CustomerID     uuid.UUID        `json:"customer_id" swaggerignore:"true"`
	Address        string           `json:"address" validate:"required,min=10,max=100" example:"123 Main St, Springfield"`
	Coords         customers.Coords `json:"coords" validate:"required"`
	MethodPayment  string           `json:"method_payment" validate:"required,oneof=Nequi Efectivo" example:"Nequi"`
	Delivery_Price int              `json:"delivery_price" validate:"gte=0" example:"5000"`
	Total_Price    int              `json:"total_price" swaggerignore:"true"`

	Items []ItemsOrder `json:"items" validate:"required,min=1"`
}

func (io *CreateOrder) Validate() error {
	_ = conform.Struct(context.Background(), io)
	return validate.Struct(io)
}
