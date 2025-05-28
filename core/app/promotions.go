package app

import (
	"CaliYa/core/domain/models"
	"CaliYa/core/domain/ports"
	"context"
)

type promotionsApp struct {
	repo ports.PromotionsRepository
}

func NewPromotionsApp(repo ports.PromotionsRepository) ports.PromotionsApp {
	return &promotionsApp{repo}
}

func (p *promotionsApp) GetPromotions(ctx context.Context) ([]models.ItemsPromo, error) {

	promos, err := p.repo.GetPromotions(ctx)
	if err != nil {
		return promos, err
	}

	return promos, nil

}
