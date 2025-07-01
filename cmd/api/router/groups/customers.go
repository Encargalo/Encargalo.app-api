package groups

import (
	"CaliYa/cmd/api/handler"

	"github.com/labstack/echo/v4"
)

type CustomersGroup interface {
	Resource(g *echo.Group)
}

type customersGroup struct {
	handlerCustomers handler.CustomersHandler
}

func NewCustomersGroup(handlerCustomers handler.CustomersHandler) CustomersGroup {
	return &customersGroup{handlerCustomers}
}

func (o *customersGroup) Resource(g *echo.Group) {
	g.POST("/customers", o.handlerCustomers.RegisterCusrtomers)
}
