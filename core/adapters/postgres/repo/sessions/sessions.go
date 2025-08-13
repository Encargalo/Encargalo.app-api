package sessions

import (
	sessionsPorts "CaliYa/core/domain/ports/sessions"
	"context"
	"database/sql"
	"errors"
	"fmt"

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

func (s *sessions) SearchSessions(ctx context.Context, id uuid.UUID) (sessionsModel.ActiveSession, error) {

	session := sessionsModel.ActiveSession{}

	if err := s.db.NewSelect().Model(&session).Where("id = ?", id).Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return sessionsModel.ActiveSession{}, errors.New("not found")
		}
		fmt.Println(err.Error())
		return sessionsModel.ActiveSession{}, errors.New("unexpected error")
	}
	return session, nil
}
