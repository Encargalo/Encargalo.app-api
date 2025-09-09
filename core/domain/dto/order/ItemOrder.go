package order

import (
	"context"

	"github.com/google/uuid"
)

type ItemsOrder struct {
	ItemID      uuid.UUID `json:"item_id" validate:"required,uuid4" example:"9ad8b85b-b847-4f15-a0ce-6415b7e335f0"`
	OrderID     uuid.UUID `swaggerignore:"true"`
	Amount      int       `json:"cant_item" validate:"required" example:"2"`
	UnitPrice   int       `swaggerignore:"true"`
	TotalPrice  int       `swaggerignore:"true"`
	Observation string    `json:"observation" example:"Con todas las salsas."`
}

func (io *ItemsOrder) Validate() error {
	_ = conform.Struct(context.Background(), io)
	return validate.Struct(io)
}
