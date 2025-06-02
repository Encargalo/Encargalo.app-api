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

type shopsRepo struct {
	db *bun.DB
}

func NewShopsRepository(db *bun.DB) ports.ShopsRepo {
	return &shopsRepo{db}
}

func (s *shopsRepo) GetAllShops(ctx context.Context) (dto.ShopsResponse, error) {

	var shops models.Shops

	if err := s.db.NewSelect().Model(&shops).Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return dto.ShopsResponse{}, echo.NewHTTPError(http.StatusNotFound, "products not found")
		}
		fmt.Println(err.Error())
		return dto.ShopsResponse{}, echo.NewHTTPError(http.StatusInternalServerError, "unexpected error")
	}

	return shops.ToDomainDTO(), nil
}
