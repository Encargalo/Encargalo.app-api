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
	middle           middleware.Request
	middleAuth       middleware.AuthMiddleware
	handlersignIn    customers.Sign_In
	handlerAddress   customers.CustomersAddressHandler
	handlerCustomers customers.CustomersHandler
}

func NewCustomersGroup(
	middle middleware.Request,
	middleAuth middleware.AuthMiddleware,
	handlersignIn customers.Sign_In,
	handlerAddress customers.CustomersAddressHandler,
	handlerCustomers customers.CustomersHandler) CustomersGroup {
	return &customersGroup{
		middle,
		middleAuth,
		handlersignIn,
		handlerAddress,
		handlerCustomers}
}

func (o *customersGroup) Resource(g *echo.Group) {
	g.POST("/customers/sign_in", o.handlersignIn.CreateSession, o.middle.GetRequestInfo)
	g.POST("/customers", o.handlerCustomers.RegisterCustomer, o.middle.GetRequestInfo)
	g.GET("/customers", o.handlerCustomers.SearchCustomer, o.middle.GetRequestInfo, o.middleAuth.Auth)
	g.POST("/customers/address", o.handlerAddress.RegisterAddress, o.middle.GetRequestInfo, o.middleAuth.Auth)
	g.GET("/customers/address", o.handlerAddress.SearchAllAdrress, o.middle.GetRequestInfo, o.middleAuth.Auth)
}
