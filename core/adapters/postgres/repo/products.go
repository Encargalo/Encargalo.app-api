package repo

import (
	"CaliYa/core/domain/dto"
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

func (p *productsRepo) GetProductsBy(ctx context.Context, criteria dto.SearchProductsBy) (*models.ProductsShops, error) {

	products := new(models.ProductsShops)

	if err := p.db.NewSelect().
		Model(products).
		Where("id = ?", criteria.ShopID).
		Relation("Categories").
		Relation("Categories.Items", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Where("is_available = ?", true).
				OrderExpr("price ASC")
		}).
		Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return products, echo.NewHTTPError(http.StatusNotFound, "products not found")
		}
		return products, echo.NewHTTPError(http.StatusInternalServerError, "unexpected error")
	}

	return products, nil
}

// func (p *productsRepo) GetProductByCategory(ctx context.Context, category string) ([]models.Items, error) {

// 	category_id := []models.Categories{}
// 	items := []models.Items{}

// 	if err := p.db.NewSelect().Model(&category_id).Where("name ILIKE ?", "%"+category+"%").
// 		Scan(ctx); err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			return []models.Items{}, echo.NewHTTPError(http.StatusNotFound, "category not found")
// 		}
// 		fmt.Println(err.Error())
// 		return []models.Items{}, echo.NewHTTPError(http.StatusInternalServerError, "unexpected error")
// 	}

// 	ids := make([]uuid.UUID, len(category_id))

// 	for i := range category_id {
// 		ids[i] = category_id[i].ID
// 	}

// 	if err := p.db.NewSelect().
// 		Model(&items).OrderExpr("price ASC").
// 		Where("category_id in (?) and is_available = ?", bun.In(ids), true).
// 		Relation("ProductsShops").
// 		Scan(ctx); err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			return []models.Items{}, echo.NewHTTPError(http.StatusNotFound, "products not found")
// 		}
// 		fmt.Println(err.Error())
// 		return []models.Items{}, echo.NewHTTPError(http.StatusInternalServerError, "unexpected error")
// 	}

// 	return items, nil
// }

func (p *productsRepo) GetProductByCategory(ctx context.Context, category string) ([]models.Items, error) {

	items := []models.Items{}

	if err := p.db.NewSelect().
		Model(&items).Join("left join products.categories AS c on c.id = items.category_id").
		OrderExpr("price ASC").
		Where("c.name ILIKE ?", "%"+category+"%").
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
