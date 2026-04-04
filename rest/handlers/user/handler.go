package user

import (
	"github.com/labib0x9/ProjectUnsafe/repo"
	middleware "github.com/labib0x9/ProjectUnsafe/rest/middleware"
)

type Handler struct {
	middlewares *middleware.Middlewares
	userRepo    repo.UserRepository
}

func NewHandler(userRepo repo.UserRepository) *Handler {
	return &Handler{
		userRepo: userRepo,
	}
}
