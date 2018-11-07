package repositories

import (
	"github.com/bombergame/profiles-service/models"
)

type ProfileRepository interface {
	Create(p models.Profile) error

	FindByID(id int64) (*models.Profile, error)
	FindByUsername(username string) (*models.Profile, error)

	Update(id int64, p models.Profile) error

	Delete(id int64) error
}
