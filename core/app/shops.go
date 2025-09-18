package app

import (
	"CaliYa/core/domain/dto"
	"CaliYa/core/domain/models/items"
	"CaliYa/core/domain/ports"
	"context"
)

type shopsApp struct {
	repo ports.ShopsRepo
}

func NewShopsApp(repo ports.ShopsRepo) ports.ShopsApp {
	return &shopsApp{repo}
}

func (s *shopsApp) GetAllShops(ctx context.Context) (dto.ShopsResponse, error) {

	shops, err := s.repo.GetAllShops(ctx)
	if err != nil {
		return make(dto.ShopsResponse, 0), err
	}

	return shops, nil
}

func (p *shopsApp) GetShopsBy(ctx context.Context, criteria dto.SearchShopsByID) (*items.ItemsShops, error) {
	return p.repo.GetShopsBy(ctx, criteria)
}
