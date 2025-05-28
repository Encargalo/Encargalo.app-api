package handler

import (
	"CaliYa/core/domain/ports"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Promos interface {
	GetPromos(c echo.Context) error
}

type promos struct {
	app ports.PromotionsApp
}

func NewPromos(app ports.PromotionsApp) Promos {
	return &promos{app}
}

func (p *promos) GetPromos(c echo.Context) error {

	ctx := c.Request().Context()

	promos, err := p.app.GetPromotions(ctx)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, promos)
}
