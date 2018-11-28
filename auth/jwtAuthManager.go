package auth

import (
	"github.com/bombergame/auth-service/config"
	"github.com/bombergame/common/auth"
	"github.com/bombergame/common/auth/jwt"
	"github.com/bombergame/common/errs"
	"strconv"
)

type JwtAuthManager struct {
	tokenManager *jwt.TokenManager
}

func NewJwtAuthManager() *JwtAuthManager {
	return &JwtAuthManager{
		tokenManager: jwt.NewTokenManager(config.TokenSignKey),
	}
}

func (m *JwtAuthManager) GetProfileInfo(authToken string, userAgent string) (*auth.ProfileInfo, error) {
	info, err := m.tokenManager.ParseToken(authToken)
	if err != nil {
		return nil, err
	}

	if info.UserAgent != userAgent {
		err := errs.NewAccessDeniedError()
		return nil, err
	}

	id, err := strconv.ParseInt(info.ProfileID, 10, 64)
	if err != nil {
		panic(err)
	}

	profileID := &auth.ProfileInfo{
		ID: id,
	}
	return profileID, nil
}
