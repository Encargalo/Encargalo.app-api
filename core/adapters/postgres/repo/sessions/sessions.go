package sessions

import (
	sessionsPorts "CaliYa/core/domain/ports/sessions"
	"context"

	sessionsModel "CaliYa/core/domain/models/sessions"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type sessions struct {
	db *bun.DB
}

func NewSessionsRepo(db *bun.DB) sessionsPorts.SessionsRepo {
	return &sessions{db}
}

func (s *sessions) RegisterSessions(ctx context.Context, session *sessionsModel.ActiveSession) (uuid.UUID, error) {

	if _, err := s.db.NewInsert().Model(session).Returning("id").Exec(ctx); err != nil {
		return uuid.Nil, err

	}

	return session.ID, nil

}
