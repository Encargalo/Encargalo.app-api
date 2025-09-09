package order

import (
	"CaliYa/core/domain/dto/order"
	"CaliYa/core/domain/ports"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
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

	order := order.CreateOrder{}

	if err := c.Bind(&order); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	}

	if err := order.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	custoID, err := uuid.Parse(strings.TrimSpace(fmt.Sprintln(ctx.Value("customer_id"))))
	if err != nil {
		fmt.Println("Error al obtener el customer_id")
		return echo.NewHTTPError(http.StatusInternalServerError, "unexpected error")
	}

	order.CustomerID = custoID

	if err := o.app.RegisterOrders(ctx, order); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, "order created success.")
}
