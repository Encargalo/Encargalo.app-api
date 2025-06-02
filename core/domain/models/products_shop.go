package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ProductsShops struct {
	bun.BaseModel `bun:"table:business.shops"`

	ID        uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Name      string    `bun:"name" json:"name"`
	Tag       string    `json:"tag"`
	LogoImage string    `bun:"logo_image" json:"logo_image"`
	Address   string    `bun:"address" json:"address"`
	Score     int       `bun:"score" json:"score"`
	Opened    bool      `bun:"opened" json:"opened"`

	Categories []Categories `bun:"rel:has-many,join:id=shop_id" json:"categories,omitempty"`
}
