package domains

import "github.com/bombergame/common/errs"

//go:generate easyjson

//easyjson:json
type Profile struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Score    int32  `json:"score"`
}

func (p Profile) Validate() error {
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
