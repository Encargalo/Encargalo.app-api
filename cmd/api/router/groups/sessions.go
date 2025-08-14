package groups

import (
	"CaliYa/cmd/api/handler/sessions"
	middleware "CaliYa/cmd/api/middleware/requets"

	"github.com/labstack/echo/v4"
)

type SessionsGroup interface {
	Resource(g *echo.Group)
}

type sessionsGroup struct {
	middle          middleware.Request
	middleAuth      middleware.AuthMiddleware
	handlerSessions sessions.SessionsHand
}

func NewSessionsGroup(
	handlerSessions sessions.SessionsHand,
	middle middleware.Request,
	middleAuth middleware.AuthMiddleware) SessionsGroup {
	return &sessionsGroup{
		middle,
		middleAuth,
		handlerSessions}
}

func (s *sessionsGroup) Resource(g *echo.Group) {
	g.DELETE("/sessions", s.handlerSessions.DeleteSession, s.middle.GetRequestInfo, s.middleAuth.Auth)
}
