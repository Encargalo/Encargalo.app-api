package repo

import (
	"CaliYa/core/domain/models"
	"CaliYa/core/domain/ports"
	"context"
	"database/sql"
	"errors"
	"net/http"

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
