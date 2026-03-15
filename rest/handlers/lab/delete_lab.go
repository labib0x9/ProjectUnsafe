package lab

import (
	"encoding/json"
	"net/http"

	"github.com/labib0x9/ProjectUnsafe/model"
)

type LabRequest struct {
	Id string
}

func (h *Handler) DeleteLab(w http.ResponseWriter, r *http.Request) {
	var newLab LabRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newLab); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	tempLabList := make([]model.Lab, 0, len(model.LabList))
	for _, lab := range model.LabList {
		if lab.Id == newLab.Id {
			continue
		}
		tempLabList = append(tempLabList, lab)
	}
	model.LabList = tempLabList
}
