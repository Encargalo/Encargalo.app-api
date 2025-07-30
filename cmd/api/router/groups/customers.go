package groups

import (
	"CaliYa/cmd/api/handler/customers"
	middleware "CaliYa/cmd/api/middleware/requets"

	"github.com/labstack/echo/v4"
)

type CustomersGroup interface {
	Resource(g *echo.Group)
}

type customersGroup struct {
	handlerCustomers customers.CustomersHandler
	handlersignIn    customers.Sign_In
	middle           middleware.Request
}

func NewCustomersGroup(handlerCustomers customers.CustomersHandler, handlersignIn customers.Sign_In, middle middleware.Request) CustomersGroup {
	return &customersGroup{handlerCustomers, handlersignIn, middle}
}

func (o *customersGroup) Resource(g *echo.Group) {
	g.POST("/customers/sign_in", o.handlersignIn.CreateSession)
	g.POST("/customers", o.handlerCustomers.RegisterCustomers, o.middle.GetRequestInfo)
}
