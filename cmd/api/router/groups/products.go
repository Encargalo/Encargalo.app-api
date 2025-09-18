package groups

import (
	productsHand "CaliYa/cmd/api/handler/products"

	"github.com/labstack/echo/v4"
)

type ProductsGroup interface {
	Resource(g *echo.Echo)
}

type productsGroup struct {
	handlerProducts productsHand.Products
	handlerFlavors  productsHand.Flavors
}

func NewProductsGroup(handlerProducts productsHand.Products, handlerFlavors productsHand.Flavors) ProductsGroup {
	return &productsGroup{handlerProducts, handlerFlavors}
}

func (r *productsGroup) Resource(g *echo.Echo) {
	g.POST("/products", r.handlerProducts.RegisterProducts)
	g.GET("/products/category", r.handlerProducts.GetProductsByCategory)
	g.GET("/products/adiciones", r.handlerProducts.GetAdicionesByCategory)
	g.GET("/products/flavors", r.handlerFlavors.SearchFlavorsBy)
}
