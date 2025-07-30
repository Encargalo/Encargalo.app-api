package customers

import (
	dto "CaliYa/core/domain/dto/customers"
	"CaliYa/core/domain/ports/customers"
	"CaliYa/core/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Sign_In interface {
	CreateSession(e echo.Context) error
}

type sign_in struct {
	svc   customers.CustomersSessionsApp
	utils utils.Sessions
}

func NewSignInCustomers(svc customers.CustomersSessionsApp, utils utils.Sessions) Sign_In {
	return &sign_in{svc, utils}
}

func (s *sign_in) CreateSession(e echo.Context) error {

	ctx := e.Request().Context()

	sign_in := dto.SignIn{}

	if err := e.Bind(&sign_in); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := sign_in.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	sessionID, err := s.svc.Sign_In(ctx, sign_in)
	if err != nil {
		if err.Error() == "incorrect access data" {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, "unexpected error")
		}
	}

	sessionJWT, err := s.utils.CreateSession(sessionID)
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

	return e.JSON(http.StatusCreated, "session created success")
}
