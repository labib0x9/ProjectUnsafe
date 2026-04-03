package user

import (
	"net/http"
	"strconv"

	"github.com/labib0x9/ProjectUnsafe/utils"
)

type ProfileResp struct {
	Username string
	Photo    string
}

func (h *Handler) GetProfile(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))
	found, err := h.userRepo.GetProfile(id)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	profile := ProfileResp{
		Username: found.Username,
		Photo:    found.ProfilePic,
	}

	utils.SendJson(w, profile, http.StatusOK)
}