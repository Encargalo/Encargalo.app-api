package handler

import (
	dto "CaliYa/core/domain/dto/customers"
	"CaliYa/core/domain/ports"
	"CaliYa/core/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type CustomersHandler interface {
	RegisterCustomers(e echo.Context) error
}

type customersHandler struct {
	customerApp ports.CustomersApp
	utils       utils.Sessions
}

func NewCustomersHandler(customerApp ports.CustomersApp, utils utils.Sessions) CustomersHandler {
	return &customersHandler{customerApp, utils}
}

// RegisterCustomers godoc
// @Summary      Registrar un nuevo cliente
// @Description  Registrar un nuevo cliente en el sistema con los datos proporcionados. Valida campos obligatorios como nombre, teléfono y contraseña.
// @Tags         Customers
// @Accept       json
// @Param        customer  body  dto.RegisterCustomer  true  "Datos del cliente"
// @Success      201  {string}  string  "customer successfully registered"
// @Failure      400 {string} string "Se retorna cuando hay un campo que no cumple con los requisitos o directamente el body se envía vacío."
// @Failure      500 {string} string "Se retorna cuando ocurre un error inexperado en el servidor."
// @Router       /customers [post]
func (c *customersHandler) RegisterCustomers(e echo.Context) error {

	ctx := e.Request().Context()

	customer := dto.RegisterCustomer{}

	if err := e.Bind(&customer); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := customer.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	sessionID, err := c.customerApp.RegisterCustomer(ctx, customer)
	if err != nil {
		switch err.Error() {
		case "phone al ready exist":
			return echo.NewHTTPError(http.StatusConflict, err.Error())
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	sessionJWT, err := c.utils.CreateSession(sessionID)
	if err != nil {
		fmt.Println(err)
	}

	cookie := &http.Cookie{
		Name:     "Sessions",
		Value:    sessionJWT,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
		HttpOnly: true,
		Path:     "/",
		Expires:  time.Now().Add(30 * 24 * time.Hour), // Expira en 30 días
	}

	e.SetCookie(cookie)

	return e.JSON(http.StatusCreated, "customer successfully registered")
}
