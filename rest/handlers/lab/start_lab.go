package lab

import (
	"log/slog"
	"net/http"
)

func (h *Handler) StartLab(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, 256)
	r.Body.Read(body)
	slog.Info(string(body))
}
