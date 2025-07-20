package app

import (
	sessionModel "CaliYa/core/domain/models/sessions"
	"CaliYa/core/domain/ports"
	"context"
	"fmt"

	"github.com/google/uuid"
)

type sessionApp struct {
	repo ports.SessionsRepo
}

func NewSessionsApp(repo ports.SessionsRepo) ports.SessionsApp {
	return &sessionApp{repo}
}

func (s *sessionApp) RegisterSessions(ctx context.Context, userID uuid.UUID, userType string) (uuid.UUID, error) {

	session := sessionModel.ActiveSession{}
	session.BuildActiveSessionModel(userID, userType, fmt.Sprint("", ctx.Value("ip")), fmt.Sprint("", ctx.Value("user-agent")))

	return s.repo.RegisterSessions(ctx, &session)

}
