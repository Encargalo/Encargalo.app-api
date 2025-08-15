package sessions

import (
	"CaliYa/core/domain/ports/sessions"
	"fmt"
	"net/http"
	"strings"
	"time"

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

// DeleteSession godoc
// @Summary Cierra la sesión actual del cliente autenticado
// @Description Elimina la sesión activa identificada por session_id en la cookie de autenticación
// @Tags Sessions
// @Produce json
// @Success 200 {string} string "session deleted success"
// @Failure 404 {string} string "not found"
// @Failure 500 {string} string "unexpected error"
// @Security SessionCookie
// @Router /sessions [delete]
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

	cookie := &http.Cookie{
		Name:     "Sessions",
		Value:    "",
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
		HttpOnly: true,
		Path:     "/",
		Expires:  time.Unix(0, 0), // Fecha en el pasado
		MaxAge:   -1,              // Forzar eliminación
	}

	e.SetCookie(cookie)

	return e.JSON(http.StatusOK, "session deleted success")
}
