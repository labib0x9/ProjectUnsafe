package lab

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/labib0x9/ProjectUnsafe/model"
	"github.com/labib0x9/ProjectUnsafe/utils"
)

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {

	if r.Header.Get("Role") != "admin" {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	var newLab model.Lab
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newLab); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	newLab, err := h.labRepo.Create(newLab)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		slog.Error("Create() failed", "error", err.Error())
		return
	}

	utils.SendJson(w, newLab, http.StatusCreated)
}
