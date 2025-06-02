package handler

import (
	"CaliYa/core/domain/models"
	"CaliYa/core/domain/ports"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Orders interface {
	RegisterOrder(c echo.Context) error
}

type orders struct {
	app ports.OrdersApp
}

func NewOrdersHandler(app ports.OrdersApp) Orders {
	return &orders{app}
}

func (o *orders) RegisterOrder(c echo.Context) error {

	ctx := c.Request().Context()

	order := models.Order{}

	if err := c.Bind(&order); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	}

	if err := o.app.RegisterOrders(ctx, order); err != nil {
		return nil
	}

	return nil
}
