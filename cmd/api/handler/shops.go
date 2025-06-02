package handler

import (
	"CaliYa/core/domain/dto"
	"CaliYa/core/domain/ports"
	calierrors "CaliYa/core/errors"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Shops interface {
	GetAllShops(c echo.Context) error
	GetShopsBy(c echo.Context) error
}

type shops struct {
	app ports.ShopsApp
}

func NewShopsHandler(app ports.ShopsApp) Shops {
	return &shops{app}
}

func (s *shops) GetAllShops(c echo.Context) error {

	ctx := c.Request().Context()

	shops, err := s.app.GetAllShops(ctx)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, shops)
}

func (p *shops) GetShopsBy(c echo.Context) error {

	ctx := c.Request().Context()

	criteria := dto.SearchShopsByID{}

	if err := c.Bind(&criteria); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := criteria.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	products, err := p.app.GetShopsBy(ctx, criteria)
	if err != nil {
		switch {
		case errors.Is(err, calierrors.ErrShopNotFound):
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
		}
	}

	return c.JSON(http.StatusOK, products)
}
