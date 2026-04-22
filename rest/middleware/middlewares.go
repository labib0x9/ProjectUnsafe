package middleware

import (
	"github.com/labib0x9/ProjectUnsafe/config"
	"github.com/labib0x9/ProjectUnsafe/repo"
)

type Middlewares struct {
	Cnf       *config.Config
	cacheRepo repo.CacheRepo
}

func NewMiddlewares(cnf *config.Config, cacheRepo repo.CacheRepo) *Middlewares {
	return &Middlewares{
		Cnf:       cnf,
		cacheRepo: cacheRepo,
	}
}
