package models

import "github.com/google/uuid"

type ProductsShops struct {
	ID         uuid.UUID    `json:"id" gorm:"primaryKey"`
	Shop       string       `json:"shop" gorm:"column:name"`
	Direction  string       `json:"direction"`
	Categories []Categories `json:"categories" gorm:"foreignKey:ProductShopID;references:ID"`
}
