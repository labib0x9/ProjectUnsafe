package lab

import (
	"github.com/labib0x9/ProjectUnsafe/repo"
	"github.com/labib0x9/ProjectUnsafe/rest/middleware"
)

type Handler struct {
	middlewares middleware.Middlewares
	labRepo     repo.LabRepository
}

func NewHandler(labRepo repo.LabRepository) *Handler {
	return &Handler{
		labRepo: labRepo,
	}
}
