package dto

import (
	"context"
)

type RegisterCustomer struct {
	Name     string `json:"name" validate:"required,min=3" example:"Carlos"`
	SurName  string `json:"sur_name" validate:"required,min=3" example:"Ramírez"`
	Phone    string `json:"phone" validate:"required,e164" example:"+573001112233"`
	Email    string `json:"email" validate:"email" example:"carlos.ramirez@example.com"`
	Password string `json:"password" validate:"required,min=8" example:"claveSegura123"`
}

func (c *RegisterCustomer) Validate() error {
	_ = conform.Struct(context.Background(), c)
	return validate.Struct(c)
}
