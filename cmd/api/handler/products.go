package handler

import (
	"CaliYa/core/domain/ports"
	calierrors "CaliYa/core/errors"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Products interface {
	RegisterProducts(c echo.Context) error
	GetProductsByCategory(c echo.Context) error
	GetAdicionesGyCategory(c echo.Context) error
}

type products struct {
	app ports.ProductsApp
}

func NewProducts(app ports.ProductsApp) Products {
	return &products{app}
}

func (p *products) RegisterProducts(c echo.Context) error {

	ctx := c.Request().Context()

	err := p.app.RegisterProduct(ctx)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "ok")
}

func (p *products) GetProductsByCategory(c echo.Context) error {

	ctx := c.Request().Context()

	category := c.Param("category")

	combos, err := p.app.GetProductByCategory(ctx, category)
	if err != nil {
		switch {
		case errors.Is(err, calierrors.ErrNotFound):
			echo.NewHTTPError(http.StatusNotFound, err)
		default:
			echo.NewHTTPError(http.StatusInternalServerError, err)
		}

	}

	return c.JSON(http.StatusOK, combos)
}

func (p *products) GetAdicionesGyCategory(c echo.Context) error {

	ctx := c.Request().Context()

	id := c.QueryParam("category_id")

	adiciones, err := p.app.GetAditionsByCategory(ctx, uuid.MustParse(id))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, adiciones)
}
