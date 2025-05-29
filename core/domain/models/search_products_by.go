package models

import (
	"context"

	"github.com/google/uuid"
)

type SearchProductsBy struct {
	ID uuid.UUID `json:"id" query:"id" validate:"uuid4"`
}

func (s *SearchProductsBy) Validate() error {
	_ = conform.Struct(context.TODO(), s)
	return validate.Struct(s)
}
