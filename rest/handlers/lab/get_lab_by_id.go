package lab

import (
	"net/http"

	"github.com/labib0x9/ProjectUnsafe/model"
	"github.com/labib0x9/ProjectUnsafe/utils"
)

func (h *Handler) GetLabByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var lab model.Lab
	found := false
	for _, tempLab := range model.LabList {
		if tempLab.Id == id {
			lab = tempLab
			found = true
			break
		}
	}

	if found == false {
		w.WriteHeader(404)
		return
	}
	utils.SendJson(w, lab)
}
