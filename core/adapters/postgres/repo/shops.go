package repo

import (
	"CaliYa/core/domain/dto"
	"CaliYa/core/domain/models"
	"CaliYa/core/domain/ports"
	calierrors "CaliYa/core/errors"
	"errors"

	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

type shopsRepo struct {
	db *bun.DB
}

func NewShopsRepository(db *bun.DB) ports.ShopsRepo {
	return &shopsRepo{db}
}

func (s *shopsRepo) GetAllShops(ctx context.Context) (dto.ShopsResponse, error) {

	var shops models.Shops

	if err := s.db.NewSelect().Model(&shops).Order("score ASC").Where("license_status = ?", "active").Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return dto.ShopsResponse{}, echo.NewHTTPError(http.StatusNotFound, "products not found")
		}
		fmt.Println(err.Error())
		return dto.ShopsResponse{}, echo.NewHTTPError(http.StatusInternalServerError, "unexpected error")
	}

	return shops.ToDomainDTO(), nil
}

func (p *shopsRepo) GetShopsBy(ctx context.Context, criteria dto.SearchShopsByID) (*models.ProductsShops, error) {

	if criteria.ID == uuid.Nil && criteria.Tag == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "at least one search criteria is required")
	}

	products := new(models.ProductsShops)

	if err := p.db.NewSelect().
		Model(products).
		WhereGroup("or", func(sq *bun.SelectQuery) *bun.SelectQuery {
			if criteria.ID != uuid.Nil {
				sq = sq.Where("id = ?", criteria.ID)
			}
			if criteria.Tag != "" {
				sq = sq.Where("tag = ?", criteria.Tag)
			}
			return sq
		}).
		Relation("Categories").
		Relation("Categories.Items", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Where("is_available = ?", true).
				OrderExpr("price ASC")
		}).
		Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, calierrors.ErrShopNotFound
		}
		return nil, fmt.Errorf("db scan failed: %w", err)
	}

	return products, nil
}
