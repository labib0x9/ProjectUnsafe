package lab

import (
	"encoding/json"
	"net/http"

	"github.com/labib0x9/ProjectUnsafe/utils"
)

type LabRequest struct {
	Id string
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {

	if r.Header.Get("Role") != "admin" {
		http.Error(w, "Bad request", http.StatusForbidden)
		return
	}

	var newLab LabRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newLab); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if err := h.labRepo.Delete(newLab.Id); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	utils.SendJson(w, "Deleted", http.StatusOK)
}
