package handler

import (
	"CaliYa/core/domain/ports"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Products interface {
	RegisterProducts(c echo.Context) error
	GetProducts(c echo.Context) error
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

func (p *products) GetProducts(c echo.Context) error {

	ctx := c.Request().Context()

	products, err := p.app.GetProducts(ctx)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, products)
}
