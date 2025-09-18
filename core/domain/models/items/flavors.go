package items

import (
	"CaliYa/core/domain/dto/products"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Flavors []Flavor

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

func (s *Flavor) ToDomainDTO() products.FlavorResponse {
	return products.FlavorResponse{
		ID:   s.ID,
		Name: s.Name,
	}
}

func (f *Flavors) ToDomainDTO() products.FlavorsResponse {
	var flavors products.FlavorsResponse

	for _, v := range *f {
		flavors.Add(v.ToDomainDTO())
	}

	return flavors
}
