package utils

import (
	"github.com/anaskhan96/go-password-encoder"
)

type PasswordEncoderImpl struct{}

func (pEnc *PasswordEncoderImpl) Encode(raw string) (hash string, salt string) {
	s, h := passwordEncoder.EncodePassword(raw, nil)
	return h, s
}

func (pEnc *PasswordEncoderImpl) Verify(raw string, hash string, salt string) bool {
	return passwordEncoder.VerifyPassword(raw, salt, hash, nil)
}
