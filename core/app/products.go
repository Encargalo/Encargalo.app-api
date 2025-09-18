package app

import (
	ItemsModels "CaliYa/core/domain/models/items"
	"CaliYa/core/domain/ports"
	"context"

	"github.com/google/uuid"
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

func (p *productsApp) GetProductByCategory(ctx context.Context, category string) ([]ItemsModels.Items, error) {
	return p.repo.GetProductByCategory(ctx, category)
}

func (p *productsApp) GetAditionsByCategory(ctx context.Context, id uuid.UUID) ([]ItemsModels.Items, error) {
	return p.repo.GetAditionsByCategory(ctx, id)
}
