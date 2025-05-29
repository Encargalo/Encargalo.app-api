package models

import (
	"time"

	"github.com/uptrace/bun"

	"github.com/google/uuid"
)

type Categories struct {
	bun.BaseModel `bun:"table:products.categories"`

	ID        uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	ShopID    uuid.UUID `bun:"shop_id" json:"shop_id"`
	Name      string    `bun:"name" json:"name"`
	CreatedAt time.Time `bun:"created_at,default:now()" json:"-"`
	UpdatedAt time.Time `bun:"updated_at,default:now()" json:"-"`
	DeletedAt time.Time `bun:"deleted_at" json:"-"`

	Items []Items `bun:"rel:has-many,join:id=category_id" json:"items"`
}

func (c *Categories) GetCategoriesID(cat []Categories) []uuid.UUID {

	ids := make([]uuid.UUID, len(cat))

	for i := range ids {
		ids[i] = cat[i].ID
	}

	return ids

}
