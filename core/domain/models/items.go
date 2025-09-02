package models

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Items struct {
	bun.BaseModel `bun:"table:products.items" swaggerignore:"true"`

	ID          uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	ShopID      uuid.UUID `bun:"shop_id" json:"shop_id"`
	CategoryID  uuid.UUID `bun:"category_id" json:"category_id"`
	Name        string    `bun:"name" json:"name"`
	Price       int       `bun:"price" json:"price"`
	Image       string    `bun:"image" json:"image"`
	Description string    `bun:"description" json:"description"`
	Score       int       `bun:"score" json:"score"`
	CreatedAt   time.Time `bun:"created_at,default:now()" json:"-"`
	UpdatedAt   time.Time `bun:"updated_at,default:now()" json:"-"`
	DeletedAt   time.Time `bun:"deleted_at" json:"-"`

	ProductsShops *ProductsShops `bun:"rel:belongs-to,join:shop_id=id" json:"shop,omitempty"`
}

type ItemsOrders struct {
	bun.BaseModel `bun:"table:business.order_items" swaggerignore:"true"`

	ItemID      uuid.UUID `bun:"item_id" json:"item_id" validate:"required,uuid4" mold:"trim" example:"9ad8b85b-b847-4f15-a0ce-6415b7e335f0"`
	OrderID     uuid.UUID `bun:"order_id" swaggerignore:"true"`
	Amount      int       `bun:"amount" json:"cant_item" validate:"required" mold:"trim" example:"2"`
	UnitPrice   int       `bun:"unit_price" swaggerignore:"true"`
	TotalPrice  int       `bun:"total_price" swaggerignore:"true"`
	Observation string    `bun:"observation" json:"observation" example:"Con todas las salsas."`
}

func (io *ItemsOrders) Validate() error {
	_ = conform.Struct(context.Background(), io)
	return validate.Struct(io)
}
