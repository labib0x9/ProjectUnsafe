package middleware

import (
	"context"
	"log/slog"
	"net/http"
	"strings"

	"github.com/labib0x9/ProjectUnsafe/utils"
)

type contextKey struct{}

var claimKey = contextKey{}

func (m *Middlewares) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := r.Header.Get("Authorization")
		if h == "" || !strings.HasPrefix(h, "Bearer") {
			slog.Error("Auth Middleware", "error", "Bad Authorization Header "+h)
			http.Error(w, "Bad Authorization Header", http.StatusBadRequest)
			return
		}

		tokenStr := strings.TrimPrefix(h, "Bearer ")
		data, ok := utils.VerifyJWT(m.Cnf.JwtSecret, tokenStr)
		if ok == false {
			http.Error(w, "Invalid token", 400)
			slog.Error("Auth middleware", "error", "Invalid token")
			return
		}

		// fmt.Println("Auth Middleware", data.Role)

		ctx := context.WithValue(r.Context(), claimKey, data)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
