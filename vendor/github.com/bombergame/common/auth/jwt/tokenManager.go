package jwt

import (
	"github.com/bombergame/common/auth"
	"github.com/bombergame/common/consts"
	"github.com/bombergame/common/errs"
	"github.com/bombergame/common/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
	"time"
)

const (
	DefaultKeyLength = 64
)

type TokenManager struct {
	key []byte
}

func NewTokenManager(key string) *TokenManager {
	randSeqGen := utils.NewRandomSequenceGenerator()

	if key == consts.EmptyString {
		key = randSeqGen.Next(DefaultKeyLength)
	}

	return &TokenManager{
		key: []byte(key),
	}
}

func (m *TokenManager) CreateToken(info auth.TokenInfo) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"profile_id":  info.ProfileID,
		"user_agent":  info.UserAgent,
		"expire_time": time.Now().Format(auth.ExpireTimeFormat),
	})
	return t.SignedString(m.key)
}

func (m *TokenManager) ParseToken(token string) (*auth.TokenInfo, error) {
	invFmtErr := errs.NewInvalidFormatError("wrong token")

	t, err := jwt.Parse(token, func(tk *jwt.Token) (interface{}, error) {
		if _, ok := tk.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, invFmtErr
		}
		return m.key, nil
	})

	if err != nil || !t.Valid {
		return nil, invFmtErr
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return nil, invFmtErr
	}

	info := &auth.TokenInfo{}

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result: info, TagName: "mapstructure",
	})
	if err != nil {
		return nil, errs.NewServiceError(err)
	}

	if err := decoder.Decode(claims); err != nil {
		return nil, invFmtErr
	}

	return info, nil
}
