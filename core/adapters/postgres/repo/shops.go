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

	if err := s.db.NewSelect().Model(&shops).Order("score DESC").
		Where("license_status = ? and opened = ? and tag != ?", "active", true, "test").
		Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return dto.ShopsResponse{}, fmt.Errorf("shops %v", calierrors.ErrNotFound)
		}
		fmt.Println(err.Error())
		return dto.ShopsResponse{}, calierrors.ErrUnexpected
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
		Relation("Categories", func(sq *bun.SelectQuery) *bun.SelectQuery {
			return sq.Where("name != ?", "Adiciones").OrderExpr(`
		CASE 
			WHEN name = 'Combos' THEN 0 
			ELSE 1
		END, name ASC`)
		}).
		Relation("Categories.Items", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Where("is_available = ?", true).
				OrderExpr("price ASC")
		}).
		Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, calierrors.ErrNotFound
		}
		fmt.Println(err.Error())
		return nil, calierrors.ErrUnexpected
	}

	return products, nil
}
