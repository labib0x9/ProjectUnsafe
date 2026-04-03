package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/labib0x9/ProjectUnsafe/config"
	_ "github.com/lib/pq"
)

func newConnectionString(cfg *config.Config) string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s", cfg.DBUser, cfg.DBPass, cfg.DBAddr, cfg.DBPort, cfg.DBName, cfg.DBSslMode)
}

func NewConnection(cfg *config.Config) (*sqlx.DB, error) {
	dbSource := newConnectionString(cfg)
	return sqlx.Connect("postgres", dbSource)
}
