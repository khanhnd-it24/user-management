package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"user-management/core/entities"
)

func GenerateTokenAndHandleError(user entities.User, minutes time.Duration) (*string, error, int) {
	secret := GoDotEnvVariable("JWT_SECRET")
	token, err := GenerateToken(user, secret, minutes)

	if err != nil {
		return nil, err, 500
	}

	return &token, err, 200

}

func GenerateToken(user entities.User, secret string, minutes time.Duration) (string, error) {
	aToken := jwt.New(jwt.SigningMethodHS256)
	claims := aToken.Claims.(jwt.MapClaims)
	claims["id"] = user.Id
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Minute * minutes).Unix()
	claims["role"] = user.Role

	token, err := aToken.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return token, err
}

func VerifyToken(token string) (*jwt.Token, error) {
	data, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(GoDotEnvVariable("JWT_SECRET")), nil
	})

	return data, err
}
