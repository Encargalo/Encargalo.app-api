package sessions

import (
	"CaliYa/core/domain/models/sessions"
	"context"

	"github.com/google/uuid"
)

type SessionsApp interface {
	RegisterSessions(ctx context.Context, userID uuid.UUID, userType string) (uuid.UUID, error)
	SearchSessions(ctx context.Context, id uuid.UUID) (sessions.ActiveSession, error)
}

type SessionsRepo interface {
	RegisterSessions(ctx context.Context, session *sessions.ActiveSession) (uuid.UUID, error)
	SearchSessions(ctx context.Context, id uuid.UUID) (sessions.ActiveSession, error)
}
