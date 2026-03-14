package lab

import (
	"net/http"

	"github.com/labib0x9/ProjectUnsafe/model"
	"github.com/labib0x9/ProjectUnsafe/utils"
)

func (h *Handler) GetAllLabs(w http.ResponseWriter, r *http.Request) {
	utils.SendJson(w, model.LabList)
}
