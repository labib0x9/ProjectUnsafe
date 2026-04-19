package auth

import (
	"database/sql"
	"errors"
	"log/slog"
	"net/http"

	"github.com/labib0x9/ProjectUnsafe/utils"
)

func (h *Handler) Verify(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token == "" {
		http.Error(w, "token is empty", http.StatusBadRequest)
		return
	}

	hash := utils.GetTokenHash(token)
	verifier, err := h.authRepo.GetTokenByHash(hash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "token expired or invalid", http.StatusGone)
			slog.Warn("Verify: token expired", "error", err, "token_hash", hash)
			return
		}
		http.Error(w, "internal server error, try again later", http.StatusInternalServerError)
		slog.Error("Verify: token fetch error", "error", err, "token_hash", hash)
		return
	}

	if err := h.authRepo.SetVerified(verifier.UserId); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		slog.Error("Verify: user verification error", "error", err, "user_id", verifier.UserId, "token_hash", hash)
		return
	}

	if err := h.authRepo.DeleteVerifier(verifier.Id); err != nil {
		slog.Error("Verify: failed to delete verifier", "error", err, "id", verifier.Id)
	}

	utils.SendJson(w, "account verified", 200)
}
