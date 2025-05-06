package models

import (
	"github.com/google/uuid"
)

type Items struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Price       int       `json:"price"`
	Ingredients []string  `json:"ingredients"`
	Sauces      []string  `json:"sauces"`
	Category    string    `json:"category"`
	Image       string    `json:"image"`
}
