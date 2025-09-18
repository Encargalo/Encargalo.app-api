package products

import (
	"CaliYa/core/domain/models/items"
	"CaliYa/core/domain/ports/products"
	"context"

	"github.com/google/uuid"
)

type flavor struct {
	repo products.FlavorsRepo
}

func NewFlavorService(repo products.FlavorsRepo) products.FlavorApp {
	return &flavor{repo: repo}
}

func (f *flavor) SearchFlavorsByProductID(ctx context.Context, product_id uuid.UUID) ([]items.Flavor, error) {
	return f.repo.SearchFlavorsByProductID(ctx, product_id)
}
