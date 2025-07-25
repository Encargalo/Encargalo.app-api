package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Address struct {
	bun.BaseModel `bun:"table:customers.address"`

	ID         uuid.UUID  `bun:"id,pk,type:uuid,default:uuid_generate_v4()"`
	CustomerID uuid.UUID  `bun:"customer_id,notnull,type:uuid"`
	Name       string     `bun:"name,notnull"`
	Address    string     `bun:"address,notnull"`
	Latitude   *float64   `bun:"latitude"`
	Longitude  *float64   `bun:"longitude"`
	CreatedAt  time.Time  `bun:"created_at,default:now()"`
	UpdatedAt  time.Time  `bun:"updated_at,default:now()"`
	DeletedAt  *time.Time `bun:"deleted_at,nullzero"`
}
