package server

import (
	"context"
	"movie-reservation-system/internal/security"
	"net/http"
	"strings"
)

type contextKey string

const userContextKey contextKey = "user"

func JWTMiddleware(next http.Handler)http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == ""{
			http.Error(w, "Missing header", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer")

		claims, err := security.ValidateToken(tokenString)

		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), userContextKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUserFromContext(r *http.Request) map[string]any {
	claims, ok := r.Context().Value(userContextKey).(map[string]any)
	if !ok {
		return nil
	}

	return claims
}

