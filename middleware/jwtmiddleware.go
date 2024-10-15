package middleware

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"go-jwt-mux/config"
	"go-jwt-mux/helper"
	"net/http"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Retrieve token from cookie
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				response := map[string]string{"message": "unauthorized"}
				helper.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			}
			response := map[string]string{"message": "bad request"}
			helper.ResponseJSON(w, http.StatusBadRequest, response)
			return
		}

		tokenString := c.Value

		// Define claims
		claims := &config.JWTClaim{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JWT_KEY, nil
		})

		if err != nil {
			// Check if the error is due to the token being expired
			if errors.Is(err, jwt.ErrTokenExpired) {
				response := map[string]string{"message": "unauthorized, token is expired"}
				helper.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			}

			// Check if the error is due to an invalid signature
			if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
				response := map[string]string{"message": "signature invalid"}
				helper.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			}

			response := map[string]string{"message": "unauthorized"}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		}

		// Check if token is valid
		if !token.Valid {
			response := map[string]string{"message": "unauthorized"}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		}

		// Pass execution to the next handler
		next.ServeHTTP(w, r)
	})
}
