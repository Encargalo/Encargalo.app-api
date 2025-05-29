package handler

import (
	"CaliYa/core/domain/models"
	"CaliYa/core/domain/ports"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Products interface {
	RegisterProducts(c echo.Context) error
	GetProductsBy(c echo.Context) error
	GetCombos(c echo.Context) error
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

func (p *products) GetProductsBy(c echo.Context) error {

	ctx := c.Request().Context()

	criteria := models.SearchProductsBy{}

	if err := c.Bind(&criteria); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := criteria.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	products, err := p.app.GetProductsBy(ctx, criteria)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, products)
}

func (p *products) GetCombos(c echo.Context) error {

	ctx := c.Request().Context()

	combos, err := p.app.GetCombos(ctx)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, combos)
}
