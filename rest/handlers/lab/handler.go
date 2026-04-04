package lab

import (
	"github.com/labib0x9/ProjectUnsafe/repo"
	"github.com/labib0x9/ProjectUnsafe/rest/middleware"
)

type Handler struct {
	middlewares *middleware.Middlewares
	labRepo     repo.LabRepository
}

func NewHandler(labRepo repo.LabRepository, middlewares *middleware.Middlewares) *Handler {
	return &Handler{
		labRepo:     labRepo,
		middlewares: middlewares,
	}
}
