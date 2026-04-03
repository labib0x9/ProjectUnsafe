package auth

import "github.com/labib0x9/ProjectUnsafe/repo"

type Handler struct {
	authRepo repo.AuthRepository
}

func NewHandler(authRepo repo.AuthRepository) *Handler {
	return &Handler{authRepo: authRepo}
}
