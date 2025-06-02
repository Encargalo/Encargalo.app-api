package handler

import "github.com/labstack/echo/v4"

type Shops interface {
	GetAllShops(c echo.Context) error
}

type shops struct{}

func NewShopsHandler() {

}
