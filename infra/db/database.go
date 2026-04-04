package db

import (
	"github.com/jmoiron/sqlx"
	"github.com/labib0x9/ProjectUnsafe/config"
)

type Database interface {
	NewConnection(cfg *config.DbConfig) (*sqlx.DB, error)
	NewSuperConnection(cfg *config.DbConfig) (*sqlx.DB, error)
	Setup(cnf *config.DbConfig) error
	SetupAndConnection(cnf *config.DbConfig) *sqlx.DB
}

