package jwtx

import "github.com/golang-jwt/jwt/v4"

type User struct {
	Uid int64 `json:"uid"`
}

type Claims struct {
	Payload any `json:"payload"`
	jwt.RegisteredClaims
}
