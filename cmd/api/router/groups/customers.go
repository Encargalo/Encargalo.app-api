package groups

import (
	"CaliYa/cmd/api/handler"
	middleware "CaliYa/cmd/api/middleware/requets"

	"github.com/labstack/echo/v4"
)

type CustomersGroup interface {
	Resource(g *echo.Group)
}

type customersGroup struct {
	handlerCustomers handler.CustomersHandler
	middle           middleware.Request
}

func NewCustomersGroup(handlerCustomers handler.CustomersHandler, middle middleware.Request) CustomersGroup {
	return &customersGroup{handlerCustomers, middle}
}

func (o *customersGroup) Resource(g *echo.Group) {
	g.POST("/customers", o.handlerCustomers.RegisterCustomers, o.middle.GetRequestInfo)
}
