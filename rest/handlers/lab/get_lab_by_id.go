package lab

import (
	"net/http"

	"github.com/labib0x9/ProjectUnsafe/utils"
)

func (h *Handler) GetLabID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	lab, err := h.labRepo.Get(id)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	utils.SendJson(w, lab, http.StatusOK)
}
