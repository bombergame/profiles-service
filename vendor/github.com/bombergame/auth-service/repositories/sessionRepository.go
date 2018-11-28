package repositories

import (
	"github.com/bombergame/auth-service/domains"
)

type SessionRepository interface {
	AddSession(session domains.Session) error
	RefreshSession(session domains.Session) error
	DeleteSession(session domains.Session) error
	DeleteAllSessions(profileID int64) error
}
