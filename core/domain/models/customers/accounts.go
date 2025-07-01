package models

import (
	"CaliYa/core/domain/dto"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Accounts struct {
	bun.BaseModel `bun:"table:customers.accounts"`

	ID               uuid.UUID  `bun:"id,pk,type:uuid,default:uuid_generate_v4()"`
	Name             string     `bun:"name,notnull"`
	SurName          string     `bun:"sur_name,notnull"`
	Phone            string     `bun:"phone,notnull"`
	Email            *string    `bun:"email"`
	Password         []byte     `bun:"password,notnull"`
	ActivationStatus string     `bun:"activation_status,default:'in progress'"`
	CreatedAt        time.Time  `bun:"created_at,default:now()"`
	UpdatedAt        time.Time  `bun:"updated_at,default:now()"`
	DeletedAt        *time.Time `bun:"deleted_at,nullzero"`

	Addresses []*Address `bun:"rel:has-many,join:id=customer_id"`
}

func (c *Accounts) BuildCustomerRegisterModel(customer dto.RegisterCustomer) {
	c.Name = customer.Name
	c.SurName = customer.SurName
	c.Phone = customer.Phone
	c.Email = &customer.Email
	c.Password = []byte(customer.Password)
}
