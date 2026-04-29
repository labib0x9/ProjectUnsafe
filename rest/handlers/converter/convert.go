package converter

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/labib0x9/ProjectUnsafe/model"
	"github.com/labib0x9/ProjectUnsafe/utils"
)

type convertRequ struct {
	Files []string `json:"files" validate:"required"`
}

func (h *Handler) Convert(w http.ResponseWriter, r *http.Request) {
	var req convertRequ
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		slog.Warn("Convert: bad json body", "error", err)
		return
	}

	if err := h.validate.Struct(req); err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		slog.Warn("Convert: struct validation failed", "error", err)
		return
	}

	// create a job and push the job
	newJob := model.Job{
		Files:  req.Files,
		Id:     utils.GenerateRandomID().String(),
		Status: "queued",
	}
	// To-do
	// Push job to worker pool

	key := "messaage_queue:job_id:" + newJob.Id
	if err := h.cacheRepo.Set(key, newJob.Status, 5*time.Minute); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		slog.Warn("Convert: message queue failed", "error", err)
		return
	}

	utils.SendJson(w, map[string]string{
		"job_id": newJob.Id,
		"status": newJob.Status,
	}, http.StatusOK)
}
