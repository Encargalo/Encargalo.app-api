package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Shops struct {
	bun.BaseModel `bun:"table:business.shops"`

	ID   uuid.UUID `bun:"id,pk" json:"id"`
	Name string    `json:"name"`
}
