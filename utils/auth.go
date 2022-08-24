package utils

import (
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
	"user-management/core/entities"
)

func GenerateTokenAndHandleError(user entities.User, minutes time.Duration) (*string, error, int) {
	token, err := GenerateToken(user, GoDotEnvVariable("JWT_SECRET"), minutes)

	if err != nil {
		return nil, err, 500
	}

	return &token, err, 200

}

func GenerateToken(user entities.User, secret string, minutes time.Duration) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"Issure":    strconv.Itoa(int(user.Id)),
		"Role":      user.Role,
		"ExpiresAt": time.Now().Add(time.Minute * minutes).Unix(),
	})

	token, err := claims.SignedString([]byte(secret))
	return token, err
}

func VerifyToken(token string) (*jwt.Token, error) {
	data, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(GoDotEnvVariable("JWT_SECRET")), nil
	})

	return data, err
}
