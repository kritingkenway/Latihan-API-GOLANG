package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtRahasia = []byte("SuperRahasiaJWT")

func GenerateJWT(userId int) (string, error) {
	claims := jwt.MapClaims{
		"user_id" : userId,
		"exp" : time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

	signedToken, err := token.SignedString(jwtRahasia)

	if err != nil {
		return "",err
	}

	return signedToken, nil
}