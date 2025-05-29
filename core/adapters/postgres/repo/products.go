package repo

import (
	"CaliYa/core/domain/models"
	"CaliYa/core/domain/ports"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

type productsRepo struct {
	db *bun.DB
}

func NewProductsRepo(db *bun.DB) ports.ProductsRepo {
	return &productsRepo{db}
}

func (p *productsRepo) GetProductsBy(ctx context.Context, criteria models.SearchProductsBy) (*models.ProductsShops, error) {

	products := new(models.ProductsShops)

	if err := p.db.NewSelect().
		Model(products).
		Where("id = ?", criteria.ShopID).
		Relation("Categories").
		Relation("Categories.Items", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.OrderExpr("price ASC")
		}).
		Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return products, echo.NewHTTPError(http.StatusNotFound, "products not found")
		}
		return products, echo.NewHTTPError(http.StatusInternalServerError, "unexpected error")
	}

	return products, nil
}

func (p *productsRepo) GetCombos(ctx context.Context) ([]models.Combos, error) {

	category_id := []models.Categories{}
	items := []models.Combos{}

	if err := p.db.NewSelect().Model(&category_id).Where("name = ?", "Combos").
		Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []models.Combos{}, echo.NewHTTPError(http.StatusNotFound, "products not found")
		}
		fmt.Println(err.Error())
		return []models.Combos{}, echo.NewHTTPError(http.StatusInternalServerError, "unexpected error")
	}

	ids := make([]uuid.UUID, len(category_id))

	for i := range category_id {
		ids[i] = category_id[i].ID
	}

	if err := p.db.NewSelect().
		Model(&items).
		Where("category_id in (?)", bun.In(ids)).
		Relation("ProductsShops").
		Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []models.Combos{}, echo.NewHTTPError(http.StatusNotFound, "products not found")
		}
		fmt.Println(err.Error())
		return []models.Combos{}, echo.NewHTTPError(http.StatusInternalServerError, "unexpected error")
	}

	return items, nil
}
