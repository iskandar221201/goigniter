package middleware

import (
	"context"
	"net/http"
	"slices"

	"github.com/iskandar221201/goigniter/system"
)

func JWTProtected(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("Authorization")
			if len(auth) < 7 || auth[:7] != "Bearer " {
				system.Error(w, http.StatusUnauthorized, "Unauthorized")
				return
			}

			token := auth[7:]
			claims, err := system.ParseToken(token, secret)
			if err != nil {
				system.Error(w, http.StatusUnauthorized, "Unauthorized")
				return
			}

			ctx := context.WithValue(r.Context(), "user", claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func RoleGuard(roles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claims, ok := r.Context().Value("user").(*system.JWTClaims)
			if !ok {
				system.Error(w, http.StatusUnauthorized, "Unauthorized")
				return
			}

			if slices.Contains(roles, claims.Role) {
				next.ServeHTTP(w, r)
				return
			}

			system.Error(w, http.StatusForbidden, "Forbidden")
		})
	}
}
