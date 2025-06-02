package app

import (
	"CaliYa/core/domain/dto"
	"CaliYa/core/domain/models"
	"CaliYa/core/domain/ports"
	"context"
)

type productsApp struct {
	repo ports.ProductsRepo
}

func NewProductsApp(repo ports.ProductsRepo) ports.ProductsApp {
	return &productsApp{repo}
}

func (p *productsApp) RegisterProduct(ctx context.Context) error {

	// if err := p.repo.RegisterProducts(ctx); err != nil {
	// 	return err
	// }

	return nil
}

func (p *productsApp) GetProductsBy(ctx context.Context, crieria dto.SearchProductsBy) (*models.ProductsShops, error) {

	products, err := p.repo.GetProductsBy(ctx, crieria)
	if err != nil {
		return products, err
	}

	return products, nil
}

func (p *productsApp) GetProductByCategory(ctx context.Context, category string) ([]models.Items, error) {

	combos, err := p.repo.GetProductByCategory(ctx, category)
	if err != nil {
		return []models.Items{}, err
	}

	return combos, nil
}
