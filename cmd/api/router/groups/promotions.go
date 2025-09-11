package groups

import (
	"CaliYa/cmd/api/handler"

	"github.com/labstack/echo/v4"
)

type PromotionsGroup interface {
	Resource(g *echo.Echo)
}

type promotionsGroup struct {
	handlerPromotions handler.Promos
}

func NewPromotionsGroup(handlerPromotions handler.Promos) PromotionsGroup {
	return &promotionsGroup{handlerPromotions}
}

func (p *promotionsGroup) Resource(g *echo.Echo) {
	g.GET("/promotions", p.handlerPromotions.GetPromos)
}
