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

// GetAllShops godoc
// @Tags Shops
// @Summary Para que no me fastidies de cual es la ruta
// @Produce json
// @Success 200 {object} []dto.ShopResponse
// @Failure 404
// @Failure 500
// @Router /shops/all [get]
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

// GetShopsBy godoc
// @Tags Shops
// @Summary End Point para obtener un negocio con todos sus products, se debe enviar alguno de los 2 query params requeridos.
// @Produce json
// @Param id query string false "Este es el ID del negocio, viene en formato UUID"
// @Param tag query string false "Este es el tag del negocio .ej:dmo"
// @Success 200 {object} models.ProductsShops
// @Failure 404
// @Failure 500
// @Router /shops [get]
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
