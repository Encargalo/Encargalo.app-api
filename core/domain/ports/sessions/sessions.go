package sessions

import (
	"CaliYa/core/domain/models/sessions"
	"context"

	"github.com/google/uuid"
)

type SessionsApp interface {
	RegisterSessions(ctx context.Context, userID uuid.UUID, userType string) (uuid.UUID, error)
	SearchSessions(ctx context.Context, id uuid.UUID) (*sessions.ActiveSession, error)
	DeleteSession(ctx context.Context, session_id uuid.UUID) error
}

type SessionsRepo interface {
	RegisterSessions(ctx context.Context, session *sessions.ActiveSession) (uuid.UUID, error)
	SearchSessions(ctx context.Context, id uuid.UUID) (*sessions.ActiveSession, error)
	DeleteSession(ctx context.Context, session *sessions.ActiveSession) error
}
