package utils

import (
	"CaliYa/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claim struct {
	SessionID uuid.UUID `json:"session_id"`
	jwt.RegisteredClaims
}

type Sessions interface {
	CreateSession(sessionID uuid.UUID) (string, error)
}

type sessions struct {
	config config.Config
}

func NewSessionUtils(config config.Config) Sessions {
	return &sessions{config}
}

func (s *sessions) CreateSession(sessionID uuid.UUID) (string, error) {

	configJWT := s.config.JWT

	claims := Claim{
		SessionID:        sessionID,
		RegisteredClaims: jwt.RegisteredClaims{},
	}

	claims.ExpiresAt = jwt.NewNumericDate(time.Now().AddDate(0, 1, 0))

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(configJWT.Secret))

}
