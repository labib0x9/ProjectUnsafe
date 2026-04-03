package lab

import (
	"encoding/json"
	"net/http"

	"github.com/labib0x9/ProjectUnsafe/model"
	"github.com/labib0x9/ProjectUnsafe/utils"
)

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {

	if r.Header.Get("Role") != "admin" {
		http.Error(w, "Bad request", http.StatusForbidden)
		return
	}

	var newLab model.Lab
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newLab); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if err := h.labRepo.Update(newLab); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	utils.SendJson(w, "Updated", http.StatusOK)
}
