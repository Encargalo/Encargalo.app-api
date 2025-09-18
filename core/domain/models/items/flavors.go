package items

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Flavor struct {
	bun.BaseModel `bun:"table:products.flavors" swaggerignore:"true"`

	ID          uuid.UUID  `bun:"id,pk,type:uuid,default:uuid_generate_v4()"`
	ProductID   uuid.UUID  `bun:"product_id,notnull,type:uuid"`
	Name        string     `bun:"name,notnull"`
	Description *string    `bun:"description,nullzero"`
	CreatedAt   time.Time  `bun:"created_at,nullzero,notnull,default:now()"`
	UpdatedAt   time.Time  `bun:"updated_at,nullzero,notnull,default:now()"`
	DeletedAt   *time.Time `bun:"deleted_at,soft_delete,nullzero"`
}
