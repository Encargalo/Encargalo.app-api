package models

import "github.com/google/uuid"

type Categories struct {
	ID            string    `json:"id" gorm:"primaryKey"`
	Name          string    `json:"name"`
	ProductShopID uuid.UUID `json:"-" gorm:"type:uuid"` // clave foránea
	Items         []Items   `json:"items"`
}
