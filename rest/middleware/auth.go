package middleware

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labib0x9/ProjectUnsafe/utils"
)

type contextKey struct{}

var claimKey = contextKey{}

func (m *Middlewares) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := r.Header.Get("Authorization")
		if h == "" || !strings.HasPrefix(h, "Bearer ") {
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			slog.Warn("Auth Middleware: Authorization header missing", "Addr", r.RemoteAddr)
			return
		}

		tokenStr := strings.TrimPrefix(h, "Bearer ")
		data, err := utils.VerifyJWT(m.Cnf.JwtSecret, tokenStr)
		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				http.Error(w, "token expired", http.StatusUnauthorized)
				slog.Warn("Auth Middleware: token expired", "Addr", r.RemoteAddr)
				return
			}
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			slog.Warn("Auth Middleware: invalid token", "Addr", r.RemoteAddr)
			return
		}

		ctx := context.WithValue(r.Context(), claimKey, data)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
