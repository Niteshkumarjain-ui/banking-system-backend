package util

import (
	"banking-system-backend/domain"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(user domain.Users) (response domain.JwtGenerate, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})
	response.Token, err = token.SignedString([]byte(Configuration.Jwt.Token))
	return
}

func ValidateJWT(tokenString string) (response domain.JwtValidate, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(Configuration.Jwt.Token), nil
	})

	if err != nil {
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		response.Claims = claims
		return
	}
	err = errors.New("Invalid Token")
	return
}
