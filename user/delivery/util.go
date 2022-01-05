package delivery

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/flip-clean/middleware"
	"github.com/flip-clean/models"
)

type MyClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

func authenticateUser(param models.UserRegistrationPayload) (token string, err error) {
	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    middleware.AppName,
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
		Username: param.Username,
	}

	tokenClaims := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	token, err = tokenClaims.SignedString(middleware.JWTKey)
	return
}
