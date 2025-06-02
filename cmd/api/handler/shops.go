package handler

import (
	"CaliYa/core/domain/ports"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Shops interface {
	GetAllShops(c echo.Context) error
}

type shops struct {
	app ports.ShopsApp
}

func NewShopsHandler(app ports.ShopsApp) Shops {
	return &shops{app}
}

func (s *shops) GetAllShops(c echo.Context) error {

	ctx := c.Request().Context()

	shops, err := s.app.GetAllShops(ctx)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, shops)
}
