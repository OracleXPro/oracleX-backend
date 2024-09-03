package jwtx

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
)

type JWTManager struct {
	AccessSecret string
	AccessExpire time.Duration
}

func NewJwtManager(c Config) *JWTManager {
	return &JWTManager{
		AccessSecret: c.AccessSecret,
		AccessExpire: time.Duration(c.AccessExpire) * time.Second,
	}
}

func (j *JWTManager) Gen(payload any) (string, error) {
	now := time.Now()
	claims := Claims{
		Payload: payload,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "",
			Subject:   "",
			Audience:  []string{""},
			ExpiresAt: jwt.NewNumericDate(now.Add(j.AccessExpire)),
			NotBefore: jwt.NewNumericDate(now),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.AccessSecret))
}

func (j *JWTManager) Verify(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.AccessSecret), nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "fail to parse jwt")
	}

	if !tokenClaims.Valid {
		return nil, errors.Wrap(err, "invalid jwt")
	}

	claims, ok := tokenClaims.Claims.(*Claims)
	if !ok {
		return nil, errors.Wrap(err, "invalid claims")
	}
	return claims, nil
}
