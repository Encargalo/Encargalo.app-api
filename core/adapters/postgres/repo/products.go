package repo

import (
	"CaliYa/core/domain/models"
	"CaliYa/core/domain/ports"
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

func NewProductsRepo(db *bun.DB) ports.ProductsRepo {
	return &productsRepo{db}
}

func (p *productsRepo) GetProductByCategory(ctx context.Context, category string) ([]models.Items, error) {

	items := []models.Items{}

	if err := p.db.NewSelect().
		Model(&items).Join("left join products.categories AS c on c.id = items.category_id").
		OrderExpr("price ASC").
		Where("c.name ILIKE ? and is_available = ?", "%"+category+"%", true).
		Relation("ProductsShops").
		Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []models.Items{}, calierrors.ErrNotFound
		}
		fmt.Println(err.Error())
		return []models.Items{}, calierrors.ErrUnexpected
	}

	return items, nil
}

func (p *productsRepo) GetAditionsByCategory(ctx context.Context, id uuid.UUID) ([]models.Items, error) {
	adiciones := []models.Items{}

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
