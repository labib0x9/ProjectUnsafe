package auth

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/labib0x9/ProjectUnsafe/rest/middleware"
	"github.com/labib0x9/ProjectUnsafe/utils"
)

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	jwt, ok := middleware.GetAuthorizationHeader(r)
	if !ok {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		slog.Warn("Logout: failed to get authorized header", "Addr", r.RemoteAddr)
		return
	}

	claims, ok := middleware.GetClaims(r)
	if !ok {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		slog.Warn("Logout: failed to get claims", "Addr", r.RemoteAddr)
		return
	}

	key := "token_blocklist:" + jwt
	expire := time.Until(claims.ExpiresAt.Time)
	if expire <= 0 {
		utils.SendJson(w, "logout", http.StatusOK)
		return
	}

	if err := h.authRepo.Set(key, "1", expire); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		slog.Warn("Logout: failed to blocklist jwt", "Addr", r.RemoteAddr)
		return
	}

	utils.SendJson(w, "logout", http.StatusOK)
}
