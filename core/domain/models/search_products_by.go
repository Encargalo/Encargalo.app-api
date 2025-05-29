package models

import (
	"context"

	"github.com/google/uuid"
)

type SearchProductsBy struct {
	ShopID uuid.UUID `json:"shop_id" query:"shop_id" validate:"uuid4"`
}

func (s *SearchProductsBy) Validate() error {
	_ = conform.Struct(context.TODO(), s)
	return validate.Struct(s)
}
