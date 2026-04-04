package auth

import (
	"encoding/json"
	"net/http"

	"github.com/labib0x9/ProjectUnsafe/utils"
)

type reqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var requ reqLogin
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requ); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	found, err := h.authRepo.GetByEmail(requ.Email)
	if err != nil {
		http.Error(w, "Auth failed", http.StatusInternalServerError)
		return
	}

	if utils.CompareHashAndPassword(found.PasswordHash, requ.Password, h.middlewares.Cnf.HashPepper) == false {
		http.Error(w, "Auth Failed", http.StatusBadRequest)
		return
	}

	claims := utils.Payload{
		Sub:       found.Username,
		FirstName: found.Fullname,
		Role:      found.Role,
	}

	token := utils.CreateJWT(h.middlewares.Cnf.JwtSecret, claims)
	if token == "" {
		http.Error(w, "Auth failed", http.StatusInternalServerError)
		return
	}

	utils.SendJson(w, token, http.StatusOK)
}
