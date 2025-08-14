package dto

import (
	"errors"

	"github.com/google/uuid"
)

type Category struct {
	ID uuid.UUID `query:"category_id"`
}

func (c *Category) IsValid() error {

	if c.ID == uuid.Nil {
		return errors.New("the value cannot be empty")
	}

	_, err := uuid.Parse(c.ID.String())
	if err != nil {
		return errors.New("invalid uuid")
	}

	return nil
}
