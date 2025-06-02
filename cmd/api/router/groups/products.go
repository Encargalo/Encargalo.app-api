package groups

import (
	"CaliYa/cmd/api/handler"

	"github.com/labstack/echo/v4"
)

type ProductsGroup interface {
	Resource(g *echo.Group)
}

type productsGroup struct {
	handlerProducts handler.Products
}

func NewProductsGroup(handlerProducts handler.Products) ProductsGroup {
	return &productsGroup{handlerProducts}
}

func (r *productsGroup) Resource(g *echo.Group) {
	g.POST("/products", r.handlerProducts.RegisterProducts)
	g.GET("/products/category/:category", r.handlerProducts.GetProductsByCategory)
}
