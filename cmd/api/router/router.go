package router

import (
	"CaliYa/cmd/api/handler"
	"CaliYa/cmd/api/router/groups"
	"CaliYa/config"
	"net/http"

	_ "CaliYa/docs"

	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	server          *echo.Echo
	productsGroup   groups.ProductsGroup
	orderGroup      groups.OrdersGroup
	promotionsGroup groups.PromotionsGroup
	shopsGroup      groups.ShopsGroup
	customersGroup  groups.CustomersGroup
	sessionsGroup   groups.SessionsGroup
	config          config.Config
}

func New(
	server *echo.Echo,
	productsGroup groups.ProductsGroup,
	orderGroup groups.OrdersGroup,
	promotionsGroup groups.PromotionsGroup,
	shopsGroup groups.ShopsGroup,
	customersGroup groups.CustomersGroup,
	sessionsGroup groups.SessionsGroup,
	config config.Config,
) *Router {
	return &Router{
		server,
		productsGroup,
		orderGroup,
		promotionsGroup,
		shopsGroup,
		customersGroup,
		sessionsGroup,
		config,
	}
}

func (r *Router) Init() {

	r.server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "remote_ip=${remote_ip}, method=${method}, uri=${uri}, status=${status}, latency=${latency_human}\n\n",
	}))

	r.server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:                             []string{r.config.Allowed.Origins},
		AllowMethods:                             []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowHeaders:                             []string{echo.HeaderContentType},
		AllowCredentials:                         true,
		UnsafeWildcardOriginWithAllowCredentials: true,
	}))

	r.server.Use(middleware.Recover())

	r.server.GET("/health", handler.HealthCheck)
	r.server.GET("/docs/*", echoSwagger.EchoWrapHandler())

	r.productsGroup.Resource(r.server)
	r.orderGroup.Resource(r.server)
	r.promotionsGroup.Resource(r.server)
	r.shopsGroup.Resource(r.server)
	r.customersGroup.Resource(r.server)
	r.sessionsGroup.Resource(r.server)
}
