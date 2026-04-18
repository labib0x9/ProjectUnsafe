package middleware

import (
	"github.com/labib0x9/ProjectUnsafe/config"
)

type Middlewares struct {
	Cnf *config.Config
}

func NewMiddlewares(cnf *config.Config) *Middlewares {
	return &Middlewares{
		Cnf: cnf,
	}
}
