package config

import "github.com/golang-jwt/jwt/v5"

type JwtCustomClaims struct {
	UserName string `json:"userName"`
	jwt.RegisteredClaims
}
