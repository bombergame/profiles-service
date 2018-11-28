package auth

import (
	"time"
)

const (
	ExpireTimeFormat = time.RFC3339

	DefaultTokenValidDuration = 15 * time.Minute
)

type TokenInfo struct {
	ProfileID  string `mapstructure:"profile_id"`
	UserAgent  string `mapstructure:"user_agent"`
	ExpireTime string `mapstructure:"expire_time"`
}

type TokenManager interface {
	CreateToken(info TokenInfo) (string, error)
	ParseToken(token string) (*TokenInfo, error)
}
