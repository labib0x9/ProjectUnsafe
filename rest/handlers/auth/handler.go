package auth

import (
	"github.com/go-playground/validator/v10"
	"github.com/labib0x9/ProjectUnsafe/repo"
	middleware "github.com/labib0x9/ProjectUnsafe/rest/middleware"
	"github.com/labib0x9/ProjectUnsafe/utils/mailer"
)

type Handler struct {
	middlewares *middleware.Middlewares
	authRepo    repo.AuthRepository
	validate    *validator.Validate
	mailer      *mailer.Mailer
}

func NewHandler(
	authRepo repo.AuthRepository,
	middlewares *middleware.Middlewares,
	validate *validator.Validate,
	mailer *mailer.Mailer,
) *Handler {
	return &Handler{
		authRepo:    authRepo,
		middlewares: middlewares,
		validate:    validate,
		mailer:      mailer,
	}
}
