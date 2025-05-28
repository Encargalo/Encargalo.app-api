package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ItemsPromo struct {
	bun.BaseModel `bun:"table:products.promotions"`

	ID         uuid.UUID `bun:"id,pk"`
	PricePromo int       `bun:"price_promo"`
	ShopID     uuid.UUID `bun:"shop_id" json:"shop_id"`
	Shops      Shops     `bun:"rel:belongs-to,join:shop_id=id" json:"shop"`
	ItemID     uuid.UUID `bun:"item_id" json:"item_id"`
	Items      Items     `bun:"rel:belongs-to,join:item_id=id" json:"items"`
}
