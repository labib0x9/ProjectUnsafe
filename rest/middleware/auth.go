package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"log/slog"
	"net/http"
	"strings"
)

func (m *Middlewares) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := r.Header.Get("Authorization")
		if h == "" {
			return
		}

		hArr := strings.Split(h, " ")
		if len(hArr) != 2 {
			return
		}

		accessToken := hArr[1]
		tokenParts := strings.Split(accessToken, ".")
		if len(tokenParts) != 3 {
			return
		}

		jwtHeader := tokenParts[0]
		jwtPayload := tokenParts[1]
		signature := tokenParts[2]

		msg := []byte(jwtHeader + "." + jwtPayload)
		newH := hmac.New(sha256.New, []byte(m.cnf.JwtSecret))
		newH.Write(msg)
		newSignature := string(newH.Sum(nil))

		slog.Info("AuthMiddleware", "newSginature=", newSignature, " ,signature=", signature)

		if newSignature != signature {
			return
		}

		next.ServeHTTP(w, r)
	})
}
