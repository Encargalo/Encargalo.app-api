package repo

import (
	"CaliYa/core/domain/models"
	"CaliYa/core/domain/ports"
	"context"
	"database/sql"
	"errors"
	"fmt"
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

func (p *productsRepo) GetProductByCategory(ctx context.Context, category string) ([]models.Items, error) {

	items := []models.Items{}

	if err := p.db.NewSelect().
		Model(&items).Join("left join products.categories AS c on c.id = items.category_id").
		OrderExpr("price ASC").
		Where("c.name ILIKE ? and is_available = ?", "%"+category+"%", true).
		Relation("ProductsShops").
		Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []models.Items{}, echo.NewHTTPError(http.StatusNotFound, "products not found")
		}
		fmt.Println(err.Error())
		return []models.Items{}, echo.NewHTTPError(http.StatusInternalServerError, "unexpected error")
	}

	return items, nil
}
