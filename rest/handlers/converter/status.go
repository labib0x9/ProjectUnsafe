package converter

import (
	"log/slog"
	"net/http"

	"github.com/labib0x9/ProjectUnsafe/utils"
)

func (h *Handler) Status(w http.ResponseWriter, r *http.Request) {
	jobId := r.PathValue("jobId")
	if jobId == "" {
		return
	}

	key := "messaage_queue:job_id:" + jobId
	val, err := h.cacheRepo.Get(key)
	if err != nil {
		// key expired..
		// return
		slog.Warn("Status: Get key failed", "Err", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	utils.SendJson(w, val, http.StatusOK)
}
