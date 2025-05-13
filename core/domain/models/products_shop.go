package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ProductsShops struct {
	bun.BaseModel `bun:"table:business.shops"`

	ID        uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Name      string    `bun:"name" json:"name"`
	LogoImage string    `bun:"logo_image" json:"logo_image"`
	HomePhone string    `bun:"home_phone" json:"home_phone"`
	Address   string    `bun:"address" json:"address"`
	Latitude  float64   `bun:"latitude" json:"latitude"`
	Longitude float64   `bun:"longitude" json:"longitude"`
	Type      string    `bun:"type" json:"type"`
	Opened    bool      `bun:"opened" json:"-"`
	CreatedAt time.Time `bun:"created_at,default:now()" json:"-"`
	UpdatedAt time.Time `bun:"updated_at,default:now()" json:"-"`
	DeletedAt time.Time `bun:"deleted_at" json:"-"`

	Categories []Categories `bun:"rel:has-many,join:id=shop_id" json:"categories"`
}
