package groups

import (
	"CaliYa/cmd/api/handler"

	"github.com/labstack/echo/v4"
)

type ShopsGroup interface {
	Resource(g *echo.Group)
}

type shopGroup struct {
	handlerShops handler.Shops
}

func NewShopsGroup(handlerShops handler.Shops) ShopsGroup {
	return &shopGroup{handlerShops}
}

func (s *shopGroup) Resource(g *echo.Group) {
	g.GET("/shops/all", s.handlerShops.GetAllShops)
	g.GET("/shops", s.handlerShops.GetShopsBy)
}
