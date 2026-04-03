package auth

import (
	"encoding/json"
	"net/http"

	"github.com/labib0x9/ProjectUnsafe/model"
	"github.com/labib0x9/ProjectUnsafe/utils"
)

func (h *Handler) Signup(w http.ResponseWriter, r *http.Request) {
	var newUser model.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newUser); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	found, err := h.authRepo.GetByEmail(newUser.Email)
	if err == nil || newUser.Email == found.Email {
		utils.SendJson(w, "email exists", http.StatusForbidden)
		return
	}
	if _, err := h.authRepo.Create(newUser); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	utils.SendJson(w, "created user", http.StatusCreated)
}
