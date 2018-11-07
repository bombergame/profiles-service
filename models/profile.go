package models

import (
	"github.com/bombergame/common/errs"
)

type Profile struct {
	ID       int64    `json:"id"`
	Username string   `json:"username"`
	Password Password `json:"-"`
	Email    string   `json:"email"`
	Score    int32    `json:"score"`
}

type Profiles []Profile

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

func (p NewProfileData) Prepare() Profile {
	return Profile{
		Username: p.Username,
		Password: Password{
			Raw: p.Password,
		},
		Email: p.Email,
	}
}

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

func (p ProfileDataUpdate) Prepare() Profile {
	return Profile{
		Username: p.Username,
		Password: Password{
			Raw: p.Password,
		},
		Email: p.Email,
	}
}

type Password struct {
	Raw  string `json:"-"`
	Hash string `json:"-"`
	Salt string `json:"-"`
}
