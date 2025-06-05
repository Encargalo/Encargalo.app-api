package dto

import "github.com/google/uuid"

type ShopsResponse []ShopResponse

type ShopResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Tag       string    `json:"tag"`
	Address   string    `json:"address"`
	HomePhone string    `json:"home_phone"`
	Logo      string    `json:"logo"`
	Opened    bool      `json:"opened"`
	Type      string    `json:"type"`
	Score     float64   `json:"score"`
}

func (s *ShopsResponse) Add(shop ShopResponse) {
	*s = append(*s, shop)
}
