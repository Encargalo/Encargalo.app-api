package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Owners struct {
	bun.BaseModel `bun:"table:users.owners"`

	ID        uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Name      string    `bun:"name" json:"name"`
	Email     string    `bun:"email" json:"email"`
	Phone     string    `bun:"phone" json:"phone"`
	CreatedAt time.Time `bun:"created_at,default:now()" json:"created_at"`
	UpdatedAt time.Time `bun:"updated_at,default:now()" json:"updated_at"`
	DeletedAt time.Time `bun:"deleted_at" json:"deleted_at"`

	Shops []Shops `bun:"rel:has-many,join:id=owner_id" json:"shops"`
}
