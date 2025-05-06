package app

import (
	"CaliYa/core/domain/models"
	"CaliYa/core/domain/ports"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type productsApp struct {
	repo ports.ProductsRepo
}

func NewProductsApp(repo ports.ProductsRepo) ports.ProductsApp {
	return &productsApp{repo}
}

func (p *productsApp) GetProducts(ctx context.Context) ([]models.ProductsShops, error) {

	products, err := p.repo.GetProducts(ctx)
	if err != nil {
		return make([]models.ProductsShops, 0), echo.NewHTTPError(http.StatusInternalServerError, "unexpected error")
	}

	return products, nil
}
