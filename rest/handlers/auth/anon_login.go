package auth

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/labib0x9/ProjectUnsafe/model"
	"github.com/labib0x9/ProjectUnsafe/utils"
)

func (h *Handler) AnonLogin(w http.ResponseWriter, r *http.Request) {
	deletedAt := time.Now().Add(30 * time.Minute)
	newUser := model.User{
		Username:   "Guest-",
		Email:      "",
		Role:       "anon",
		IsVerified: true,
		DeletedAt:  &deletedAt,
	}
	newUser.Username += utils.GenerateRandomID().String()
	newUser.Fullname = newUser.Username
	newUser.Email = newUser.Username + "@gmail.com"

	_, err := h.authRepo.Create(newUser)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		slog.Error("Create() failed", "error", err.Error())
		return
	}

	payload := utils.Payload{
		Sub:       newUser.Username,
		FirstName: newUser.Fullname,
		Role:      newUser.Role,
	}
	
	token := utils.CreateJWT(h.middlewares.Cnf.JwtSecret, payload)
	if token == "" {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		slog.Error("JWT Token create failed", "error", "empty token")
		return
	}

	utils.SendJson(w, map[string]any{
		"token":    token,
		"username": newUser.Username,
	}, http.StatusCreated)
}
