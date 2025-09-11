package order

import "github.com/google/uuid"

type AdditionsOrder struct {
	ItemID      uuid.UUID `json:"item_id"`
	Amount      int       `json:"cant_item" validate:"required" example:"2"`
	UnitPrice   int       `swaggerignore:"true"`
	TotalPrice  int       `swaggerignore:"true"`
	Observation string    `json:"observation" example:"Con todas las salsas."`
}
