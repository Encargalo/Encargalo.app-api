package handler

import (
	"CaliYa/core/domain/dto"
	"CaliYa/core/domain/ports"
	calierror "CaliYa/core/errors"
	"fmt"

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
		switch {
		case errors.Is(err, fmt.Errorf("shops %v", calierror.ErrNotFound)):
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, calierror.ErrUnexpected)
		}
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
		case errors.Is(err, calierror.ErrNotFound):
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, calierror.ErrUnexpected)
		}
	}

	return c.JSON(http.StatusOK, products)
}
