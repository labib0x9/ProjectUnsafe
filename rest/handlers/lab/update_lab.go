package lab

import (
	"encoding/json"
	"net/http"

	"github.com/labib0x9/ProjectUnsafe/model"
)

func (h *Handler) UpdateLab(w http.ResponseWriter, r *http.Request) {

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

	for idx, lab := range model.LabList {
		if lab.Id == newLab.Id {
			model.LabList[idx] = newLab
			break
		}
	}
}
