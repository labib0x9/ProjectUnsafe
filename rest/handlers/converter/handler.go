package converter

import (
	"github.com/go-playground/validator/v10"
	"github.com/labib0x9/ProjectUnsafe/repo"
	middleware "github.com/labib0x9/ProjectUnsafe/rest/middleware"
)

type Handler struct {
	middlewares *middleware.Middlewares
	cacheRepo   repo.CacheRepo
	validate    *validator.Validate
}

func NewHandler(cacheRepo repo.CacheRepo, validate *validator.Validate, middlewares *middleware.Middlewares) *Handler {
	return &Handler{
		cacheRepo:   cacheRepo,
		middlewares: middlewares,
		validate:    validate,
	}
}
