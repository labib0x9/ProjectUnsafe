package admin

import (
	"github.com/labib0x9/ProjectUnsafe/repo"
	"github.com/labib0x9/ProjectUnsafe/rest/middleware"
)

type Handler struct {
	middlewares *middleware.Middlewares
	adminRepo   repo.AdminRepository
}

func NewHandler(adminRepo repo.AdminRepository) *Handler {
	return &Handler{adminRepo: adminRepo}
}
