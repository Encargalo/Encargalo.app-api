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

type promotionsRepository struct {
	db *bun.DB
}

func NewPromotionsRepository(db *bun.DB) ports.PromotionsRepository {
	return &promotionsRepository{db}
}

func (p promotionsRepository) GetPromotions(ctx context.Context) ([]models.ItemsPromo, error) {

	var promotions []models.ItemsPromo

	if err := p.db.NewSelect().Model(&promotions).
		Relation("Items").
		Relation("Shops").
		Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("No rows found")
			return promotions, echo.NewHTTPError(http.StatusNotFound, "not found")
		}

		return promotions, echo.NewHTTPError(http.StatusInternalServerError, "unexpected error")

	}

	return promotions, nil
}
