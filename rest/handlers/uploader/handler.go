package uploader

import (
	"github.com/go-playground/validator/v10"
	"github.com/labib0x9/ProjectUnsafe/repo"
	middleware "github.com/labib0x9/ProjectUnsafe/rest/middleware"
)

type Handler struct {
	middlewares  *middleware.Middlewares
	uploaderRepo repo.UploaderRepository
	validate     *validator.Validate
}

func NewHandler(uploaderRepo repo.UploaderRepository, validate *validator.Validate, middlewares *middleware.Middlewares) *Handler {
	return &Handler{
		uploaderRepo: uploaderRepo,
		middlewares:  middlewares,
		validate:     validate,
	}
}
