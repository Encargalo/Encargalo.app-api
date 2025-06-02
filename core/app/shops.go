package app

import (
	"CaliYa/core/domain/dto"
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
