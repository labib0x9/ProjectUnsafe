package auth

import (
	"github.com/labib0x9/ProjectUnsafe/repo"
	middleware "github.com/labib0x9/ProjectUnsafe/rest/middleware"
)

type Handler struct {
	middlewares *middleware.Middlewares
	authRepo    repo.AuthRepository
}

func NewHandler(authRepo repo.AuthRepository, middlewares *middleware.Middlewares) *Handler {
	return &Handler{
		authRepo:    authRepo,
		middlewares: middlewares,
	}
}
