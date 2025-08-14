package middleware

import (
	"context"

	"github.com/labstack/echo/v4"
)

type Request interface {
	GetRequestInfo(next echo.HandlerFunc) echo.HandlerFunc
}

type request struct {
}

func NewRequestMiddleware() Request {
	return &request{}
}

func (r *request) GetRequestInfo(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {

		ctx := c.Request().Context()

		ip := c.RealIP()
		userAgent := c.Request().Header.Get("User-Agent")

		ctx = context.WithValue(ctx, "ip", ip)
		ctx = context.WithValue(ctx, "user-agent", userAgent)

		c.SetRequest(c.Request().WithContext(ctx))

		return next(c)
	}

}
