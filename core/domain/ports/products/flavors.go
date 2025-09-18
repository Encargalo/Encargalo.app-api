package products

import (
	"CaliYa/core/domain/models/items"
	"context"

	"github.com/google/uuid"
)

type FlavorApp interface {
	SearchFlavorsByProductID(ctx context.Context, product_id uuid.UUID) ([]items.Flavor, error)
}

type FlavorsRepo interface {
	SearchFlavorsByProductID(ctx context.Context, product_id uuid.UUID) ([]items.Flavor, error)
}
