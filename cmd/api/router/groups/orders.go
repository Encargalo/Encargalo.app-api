package groups

import (
	"CaliYa/cmd/api/handler"

	"github.com/labstack/echo/v4"
)

type OrdersGroup interface {
	Resource(g *echo.Group)
}

type ordersGroup struct {
	handlerOrders handler.Orders
}

func NewOrdersGroup(handlerOrders handler.Orders) OrdersGroup {
	return &ordersGroup{handlerOrders}
}

func (o *ordersGroup) Resource(g *echo.Group) {
	g.POST("/order", o.handlerOrders.RegisterOrder)
}
