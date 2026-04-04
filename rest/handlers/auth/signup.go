package auth

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/labib0x9/ProjectUnsafe/model"
	"github.com/labib0x9/ProjectUnsafe/utils"
)

type reqSignup struct {
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (h *Handler) Signup(w http.ResponseWriter, r *http.Request) {
	var req reqSignup
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	found, err := h.authRepo.GetByEmail(req.Email)
	if err == nil || req.Email == found.Email {
		utils.SendJson(w, "email exists", http.StatusForbidden)
		return
	}

	passHash, err := utils.GenerateHash(req.Password, h.middlewares.Cnf.HashPepper, h.middlewares.Cnf.BcryptCost)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		slog.Error("Signup()", "hash generation failed", err.Error())
		return
	}

	newUser := model.User{
		Username:     req.Username,
		Fullname:     req.Fullname,
		Email:        req.Email,
		PasswordHash: passHash,
		Role:         req.Role,
		IsVerified:   func(role string) bool { return role == "admin" }(req.Role),
	}

	if _, err := h.authRepo.Create(newUser); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		slog.Error("Signup()", "db operation failed", err.Error())
		return
	}
	utils.SendJson(w, "created user", http.StatusCreated)
}
