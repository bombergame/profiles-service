package profilesrest

//go:generate easyjson

import (
	"github.com/bombergame/common/errs"
	"github.com/bombergame/profiles-service/domains"
)

// easyjson:json
type NewProfileData struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (p NewProfileData) Validate() error {
	if p.Username == "" {
		return errs.NewInvalidFormatError("empty username")
	}
	if p.Password == "" {
		return errs.NewInvalidFormatError("empty password")
	}
	if p.Email == "" {
		return errs.NewInvalidFormatError("empty email")
	}
	return nil
}

func (p NewProfileData) Prepare() domains.Profile {
	return domains.Profile{
		Username: p.Username,
		Password: p.Password,
		Email:    p.Email,
	}
}

// easyjson:json
type ProfileDataUpdate struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (p ProfileDataUpdate) Validate() error {
	if p.Username == "" && p.Password == "" && p.Email == "" {
		return errs.NewInvalidFormatError("empty update data")
	}
	return nil
}

func (p ProfileDataUpdate) Prepare() domains.Profile {
	return domains.Profile{
		Username: p.Username,
		Password: p.Password,
		Email:    p.Email,
	}
}
