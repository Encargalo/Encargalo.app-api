package groups

import (
	"CaliYa/cmd/api/handler"

	"github.com/labstack/echo/v4"
)

type ProductsGroup interface {
	Resource(g *echo.Echo)
}

type productsGroup struct {
	handlerProducts handler.Products
}

func NewProductsGroup(handlerProducts handler.Products) ProductsGroup {
	return &productsGroup{handlerProducts}
}

func (r *productsGroup) Resource(g *echo.Echo) {
	g.POST("/products", r.handlerProducts.RegisterProducts)
	g.GET("/products/category", r.handlerProducts.GetProductsByCategory)
	g.GET("/products/adiciones", r.handlerProducts.GetAdicionesByCategory)
}
