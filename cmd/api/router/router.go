package router

import (
	"CaliYa/cmd/api/handler"
	"CaliYa/cmd/api/router/groups"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	server        *echo.Echo
	productsGroup groups.ProductsGroup
}

func New(server *echo.Echo, productsGroup groups.ProductsGroup) *Router {
	return &Router{
		server,
		productsGroup,
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

	r.productsGroup.Resource(basePath)

}
