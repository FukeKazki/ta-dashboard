package config

import "github.com/golang-jwt/jwt/v5"

type JwtCustomClaims struct {
	UserId string `json:"userId"`
	jwt.RegisteredClaims
}
