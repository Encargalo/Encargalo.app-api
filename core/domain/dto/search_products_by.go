package dto

import (
	"context"

	"github.com/go-playground/mold/modifiers"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var (
	validate = validator.New()
	conform  = modifiers.New()
)

type SearchProductsBy struct {
	ShopID uuid.UUID `json:"shop_id" query:"shop_id" validate:"uuid4"`
}

func (s *SearchProductsBy) Validate() error {
	_ = conform.Struct(context.TODO(), s)
	return validate.Struct(s)
}
