package domain

import "github.com/dgrijalva/jwt-go"

type JwtGenerate struct {
	Token string
}

type JwtValidate struct {
	Claims jwt.MapClaims
}
