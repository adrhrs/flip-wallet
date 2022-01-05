package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

var (
	JWTKey      = []byte("secret_key")
	AppName     = "flip-ewallet"
	UsernameKey = "username"
)

func AuthUser(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				w.WriteHeader(http.StatusUnauthorized)
				return nil, nil
			} else if method != jwt.SigningMethodHS256 {
				w.WriteHeader(http.StatusUnauthorized)
				return nil, nil
			}
			return JWTKey, nil
		})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(req.Context(), UsernameKey, claims[UsernameKey])
		f(w, req.WithContext(ctx))
	}

}
