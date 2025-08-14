package sessions

import (
	"CaliYa/core/domain/ports/sessions"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type SessionsHand interface {
	DeleteSession(e echo.Context) error
}

type sessionshand struct {
	svc sessions.SessionsApp
}

func NewSessionsHandler(svc sessions.SessionsApp) SessionsHand {

	return &sessionshand{svc}
}

func (s *sessionshand) DeleteSession(e echo.Context) error {

	ctx := e.Request().Context()

	session_id, err := uuid.Parse(strings.TrimSpace(fmt.Sprintln(ctx.Value("session_id"))))
	if err != nil {
		fmt.Println("Error al obtener la session_id")
		return echo.NewHTTPError(http.StatusInternalServerError, "unexpected error")
	}

	if err := s.svc.DeleteSession(ctx, session_id); err != nil {
		switch err.Error() {
		case "not found":
			return echo.NewHTTPError(http.StatusNotFound, "not found")
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, "unexpected error")
		}
	}

	return e.JSON(http.StatusOK, "session deleted success")
}
