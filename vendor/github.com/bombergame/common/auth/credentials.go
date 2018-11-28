package auth

//go:generate easyjson

import (
	"github.com/bombergame/common/consts"
	"github.com/bombergame/common/errs"
)

//easyjson:json
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c Credentials) Validate() error {
	if c.Username == consts.EmptyString {
		return errs.NewInvalidFormatError("empty username")
	}
	if c.Password == consts.EmptyString {
		return errs.NewInvalidFormatError("empty password")
	}
	return nil
}
