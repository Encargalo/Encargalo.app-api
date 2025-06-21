package models

import (
	"github.com/uptrace/bun"

	"github.com/google/uuid"
)

type Categories struct {
	bun.BaseModel `bun:"table:products.categories" swaggerignore:"true"`

	ID     uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	ShopID uuid.UUID `bun:"shop_id" json:"shop_id"`
	Name   string    `bun:"name" json:"name"`

	Items []Items `bun:"rel:has-many,join:id=category_id" json:"items"`
}
