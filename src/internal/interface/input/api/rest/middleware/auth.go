package middleware

import (
	"context"
	"movie-crud-application/src/pkg"
	"net/http"
)

func Authenticate(jwtKey string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("accessToken")
			if err != nil {
				http.Error(w, "Missing access token", http.StatusUnauthorized)
				return
			}

			claims, err := pkg.ValidateJWT(cookie.Value, jwtKey)
			if err != nil {
				http.Error(w, "Access token invalid", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), "user", claims.Uid)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
