package groups

import (
	"CaliYa/cmd/api/handler/order"

	"github.com/labstack/echo/v4"
)

type OrdersGroup interface {
	Resource(g *echo.Group)
}

type ordersGroup struct {
	handlerOrders order.Orders
}

func NewOrdersGroup(handlerOrders order.Orders) OrdersGroup {
	return &ordersGroup{handlerOrders}
}

func (o *ordersGroup) Resource(g *echo.Group) {
	g.POST("/order", o.handlerOrders.RegisterOrder)
}
