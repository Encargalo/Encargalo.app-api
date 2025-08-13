package models

import (
	dto "CaliYa/core/domain/dto/customers"
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
	Password         string     `bun:"password,notnull"`
	BirthdayDate     string     `bun:"birthday_date"`
	ActivationStatus string     `bun:"activation_status,default:'in progress'"`
	CreatedAt        time.Time  `bun:"created_at,default:now()"`
	UpdatedAt        time.Time  `bun:"updated_at,default:now()"`
	DeletedAt        *time.Time `bun:"deleted_at,nullzero"`
}

func (c *Accounts) BuildCustomerRegisterModel(customer dto.RegisterCustomer) {
	c.Name = customer.Name
	c.SurName = customer.SurName
	c.Phone = customer.Phone
	c.Email = &customer.Email
	c.BirthdayDate = customer.BirthdayDate
	c.Password = customer.Password
}

func (c *Accounts) ToDomainDTO() dto.CustomerResponse {
	return dto.CustomerResponse{
		Name:         c.Name,
		SurName:      c.SurName,
		Phone:        c.Phone,
		Email:        *c.Email,
		BirthdayDate: c.BirthdayDate,
	}
}
