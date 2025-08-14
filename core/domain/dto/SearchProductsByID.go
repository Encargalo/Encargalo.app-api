package dto

import (
	"errors"
)

type SearchProductsByCategory struct {
	Category string `query:"category"`
}

func (s *SearchProductsByCategory) IsValid() error {

	if s.Category == "" {
		return errors.New("the value cannot be empty")
	}

	if len(s.Category) < 3 {
		return errors.New("the value must be at least 3 characters long")
	}

	return nil
}
