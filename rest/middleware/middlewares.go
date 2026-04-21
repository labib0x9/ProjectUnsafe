package middleware

import (
	"github.com/labib0x9/ProjectUnsafe/config"
	"github.com/labib0x9/ProjectUnsafe/repo"
)

type Middlewares struct {
	Cnf      *config.Config
	authRepo repo.AuthRepository
}

func NewMiddlewares(cnf *config.Config, authRepo repo.AuthRepository) *Middlewares {
	return &Middlewares{
		Cnf:      cnf,
		authRepo: authRepo,
	}
}
