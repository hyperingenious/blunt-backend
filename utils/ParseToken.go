package utils

import (
	"blunt-backend/models"

	"github.com/dgrijalva/jwt-go"
)

func ParseToken(tokenString string) (claim *models.Claims, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("my_secrete_key"), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*models.Claims)

	if !ok {
		return nil, err
	}

	return claims, nil
}


