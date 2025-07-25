package router

import (
	"CaliYa/cmd/api/handler"
	"CaliYa/cmd/api/router/groups"
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
}

func New(
	server *echo.Echo,
	productsGroup groups.ProductsGroup,
	orderGroup groups.OrdersGroup,
	promotionsGroup groups.PromotionsGroup,
	shopsGroup groups.ShopsGroup,
	customersGroup groups.CustomersGroup,
) *Router {
	return &Router{
		server,
		productsGroup,
		orderGroup,
		promotionsGroup,
		shopsGroup,
		customersGroup,
	}
}

func (r *Router) Init() {

	r.server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "remote_ip=${remote_ip}, method=${method}, uri=${uri}, status=${status}, latency=${latency_human}\n\n",
	}))

	r.server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:                             []string{"*"},
		AllowMethods:                             []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowHeaders:                             []string{echo.HeaderContentType},
		AllowCredentials:                         true,
		UnsafeWildcardOriginWithAllowCredentials: true,
	}))

	r.server.Use(middleware.Recover())

	basePath := r.server.Group("/api") //customize your basePath
	basePath.GET("/health", handler.HealthCheck)
	basePath.GET("/docs/*", echoSwagger.EchoWrapHandler())

	r.productsGroup.Resource(basePath)
	r.orderGroup.Resource(basePath)
	r.promotionsGroup.Resource(basePath)
	r.shopsGroup.Resource(basePath)
	r.customersGroup.Resource(basePath)
}
