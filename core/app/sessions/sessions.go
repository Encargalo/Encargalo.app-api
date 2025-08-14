package sessions

import (
	"CaliYa/core/domain/models/sessions"
	sessionModel "CaliYa/core/domain/models/sessions"
	ports "CaliYa/core/domain/ports/sessions"
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

func (s *sessionApp) SearchSessions(ctx context.Context, id uuid.UUID) (*sessions.ActiveSession, error) {
	return s.repo.SearchSessions(ctx, id)
}

func (s *sessionApp) DeleteSession(ctx context.Context, session_id uuid.UUID) error {

	session, err := s.SearchSessions(ctx, session_id)
	if err != nil {
		return err
	}

	return s.repo.DeleteSession(ctx, session)
}
