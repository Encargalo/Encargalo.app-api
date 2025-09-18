package items

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ItemsShops struct {
	bun.BaseModel `bun:"table:business.shops" swaggerignore:"true"`

	ID        uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Name      string    `bun:"name" json:"name"`
	Tag       string    `json:"tag"`
	LogoImage string    `bun:"logo_image" json:"logo_image"`
	Banner    string    `bun:"banner" json:"banner"`
	Address   string    `bun:"address" json:"address"`
	Phone     string    `bun:"home_phone" json:"phone"`
	Score     float32   `bun:"score" json:"score"`
	Opened    bool      `bun:"opened" json:"opened"`

	Categories []Categories `bun:"rel:has-many,join:id=shop_id" json:"categories,omitempty"`
}
