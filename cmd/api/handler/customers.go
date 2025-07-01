package handler

import (
	"CaliYa/core/domain/dto"
	"CaliYa/core/domain/ports"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CustomersHandler interface {
	RegisterCusrtomers(e echo.Context) error
}

type customersHandler struct {
	customerApp ports.CustomersApp
}

func NewCustomersHandler(customerApp ports.CustomersApp) CustomersHandler {
	return &customersHandler{customerApp}
}

// RegisterCustomers godoc
// @Summary      Registrar un nuevo cliente
// @Description  Registrar un nuevo cliente en el sistema con los datos proporcionados. Valida campos obligatorios como nombre, teléfono y contraseña.
// @Tags         Customers
// @Accept       json
// @Produce      json
// @Param        customer  body  dto.RegisterCustomer  true  "Datos del cliente"
// @Success      201  {string}  string  "customer successfully registered"
// @Router       /customers [post]
func (c *customersHandler) RegisterCusrtomers(e echo.Context) error {

	ctx := e.Request().Context()

	customer := dto.RegisterCustomer{}

	if err := e.Bind(&customer); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := customer.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.customerApp.RegisterCustomer(ctx, customer); err != nil {
		return err
	}

	return e.JSON(http.StatusCreated, "customer successfully registered")
}
