package dto

import "github.com/google/uuid"

type ShopsResponse []ShopResponse

type ShopResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Address   *string   `json:"address,omitempty"`
	HomePhone string    `json:"home_phone"`
	LogoImage *string   `json:"logo_image,omitempty"`
	Opened    bool      `json:"opened"`
	Type      string    `json:"type"`
	Score     int       `json:"score"`
}

func (s *ShopsResponse) Add(shop ShopResponse) {
	*s = append(*s, shop)
}
