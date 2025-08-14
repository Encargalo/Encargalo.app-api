package middleware

import (
	"CaliYa/core/domain/ports/sessions"
	"CaliYa/core/utils"
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type AuthMiddleware interface {
	Auth(next echo.HandlerFunc) echo.HandlerFunc
}

type auth struct {
	jwt utils.Sessions
	svc sessions.SessionsApp
}

func NewAuthMidlleware(jwt utils.Sessions, svc sessions.SessionsApp) AuthMiddleware {
	return &auth{jwt, svc}
}

func (a *auth) Auth(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {

		ctx := c.Request().Context()

		cookie, err := c.Cookie("Sessions")
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "token not found")

		}

		claims, err := a.jwt.ValidateToken(cookie.Value)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "token invalid")
		}

		session_id, err := uuid.Parse(strings.TrimSpace(fmt.Sprintln(claims["session_id"])))
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}

		session, err := a.svc.SearchSessions(ctx, session_id)
		if err != nil {
			if err.Error() == "not found" {
				return echo.NewHTTPError(http.StatusUnauthorized, "session invalid")
			}
			return echo.NewHTTPError(http.StatusInternalServerError, "unexpected error")
		}

		if session.ExpiresAt.Before(time.Now()) {
			return echo.NewHTTPError(http.StatusUnauthorized, "session invalid")
		}

		ctx = context.WithValue(ctx, "customer_id", session.UserID)
		ctx = context.WithValue(ctx, "session_id", session_id)

		c.SetRequest(c.Request().WithContext(ctx))

		return next(c)

	}

}
