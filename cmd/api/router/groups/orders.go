package groups

import (
	"CaliYa/cmd/api/handler/order"
	middleware "CaliYa/cmd/api/middleware/requets"

	"github.com/labstack/echo/v4"
)

type OrdersGroup interface {
	Resource(g *echo.Group)
}

type ordersGroup struct {
	handlerOrders order.Orders
	middleAuth    middleware.AuthMiddleware
	middleReqInfo middleware.Request
}

func NewOrdersGroup(handlerOrders order.Orders, middleAuth middleware.AuthMiddleware, middleReqInfo middleware.Request) OrdersGroup {
	return &ordersGroup{handlerOrders, middleAuth, middleReqInfo}
}

func (o *ordersGroup) Resource(g *echo.Group) {
	g.POST("/order", o.handlerOrders.RegisterOrder, o.middleAuth.Auth, o.middleReqInfo.GetRequestInfo)
}
