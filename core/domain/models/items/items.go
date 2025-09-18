package items

import (
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
	IsAvailable bool      `bun:"is_available"`
	Score       float32   `bun:"score" json:"score"`
	CreatedAt   time.Time `bun:"created_at,default:now()" json:"-"`
	UpdatedAt   time.Time `bun:"updated_at,default:now()" json:"-"`
	DeletedAt   time.Time `bun:"deleted_at" json:"-"`

	Shops *ItemsShops `bun:"rel:belongs-to,join:shop_id=id" json:"shop,omitempty" swaggerignore:"true"`
}
