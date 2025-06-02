package handler

import (
	"CaliYa/core/domain/dto"
	"CaliYa/core/domain/ports"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Products interface {
	RegisterProducts(c echo.Context) error
	GetProductsBy(c echo.Context) error
	GetProductsByCategory(c echo.Context) error
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

	criteria := dto.SearchProductsBy{}

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

func (p *products) GetProductsByCategory(c echo.Context) error {

	ctx := c.Request().Context()

	category := c.Param("category")

	fmt.Println("La categoria es:", category)

	combos, err := p.app.GetProductByCategory(ctx, category)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, combos)
}
