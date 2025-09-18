package repo

import (
	itemsModels "CaliYa/core/domain/models/items"
	productsPort "CaliYa/core/domain/ports/products"
	calierrors "CaliYa/core/errors"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type productsRepo struct {
	db *bun.DB
}

func NewProductsRepo(db *bun.DB) productsPort.ProductsRepo {
	return &productsRepo{db}
}

func (p *productsRepo) GetProductByCategory(ctx context.Context, category string) ([]itemsModels.Items, error) {

	items := []itemsModels.Items{}

	if err := p.db.NewSelect().
		Model(&items).Join("left join products.categories AS c on c.id = items.category_id").
		OrderExpr("price ASC").
		Where("c.name ILIKE ? and is_available = ?", "%"+category+"%", true).
		Relation("Shops").
		Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []itemsModels.Items{}, calierrors.ErrNotFound
		}
		fmt.Println(err.Error())
		return []itemsModels.Items{}, calierrors.ErrUnexpected
	}

	return items, nil
}

func (p *productsRepo) GetAditionsByCategory(ctx context.Context, id uuid.UUID) ([]itemsModels.Items, error) {
	adiciones := []itemsModels.Items{}

	err := p.db.NewSelect().Model(&adiciones).
		Join("left join products.categories_adiciones as ca on items.id = ca.item_id").
		Where("ca.category_id = ?", id).
		Scan(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return nil, calierrors.ErrUnexpected
	}

	fmt.Println("Evalua el error")
	// Revisar si el resultado está vacío
	if len(adiciones) == 0 {
		fmt.Println("Retorna el error.")
		return nil, calierrors.ErrNotFound
	}

	fmt.Println("No retorna el error.")

	return adiciones, nil
}
