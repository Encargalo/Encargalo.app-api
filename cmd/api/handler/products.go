package handler

import (
	"CaliYa/core/domain/ports"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Products interface {
	GetProducts(c echo.Context) error
}

type products struct {
	app ports.ProductsApp
}

func NewProducts(app ports.ProductsApp) Products {
	return &products{app}
}

func (p *products) GetProducts(c echo.Context) error {

	fmt.Println("mensaje desde el handler")

	ctx := c.Request().Context()

	products, err := p.app.GetProducts(ctx)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, products)
}
