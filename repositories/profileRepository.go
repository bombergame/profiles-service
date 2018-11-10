package repositories

import (
	"github.com/bombergame/profiles-service/domains"
)

type ProfileRepository interface {
	Create(p domains.Profile) error

	FindByID(id int64) (*domains.Profile, error)
	FindByUsername(username string) (*domains.Profile, error)

	Update(id int64, p domains.Profile) error

	Delete(id int64) error
}
