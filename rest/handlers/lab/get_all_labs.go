package lab

import (
	"net/http"

	"github.com/labib0x9/ProjectUnsafe/utils"
)

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	list, err := h.labRepo.List()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	utils.SendJson(w, list, http.StatusOK)
}
