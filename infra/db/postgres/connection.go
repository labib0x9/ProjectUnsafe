package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/labib0x9/ProjectUnsafe/config"
	_ "github.com/lib/pq"
)

func newConnectionString(cfg *config.DbConfig) string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s", cfg.DBUser, cfg.DBPass, cfg.DBAddr, cfg.DBPort, cfg.DBName, cfg.DBSslMode)
}

func (p *PostgreSQL) NewConnection(cfg *config.DbConfig) (*sqlx.DB, error) {
	dbSource := newConnectionString(cfg)
	return sqlx.Connect("postgres", dbSource)
}

func newSuperConnectionString(cfg *config.DbConfig) string {
	return fmt.Sprintf("user=%s password= host=%s port=%s dbname=%s sslmode=%s", cfg.DBSuperUser, cfg.DBAddr, cfg.DBPort, cfg.DBSuperDB, cfg.DBSslMode)
}

func (p *PostgreSQL) NewSuperConnection(cfg *config.DbConfig) (*sqlx.DB, error) {
	dbSource := newSuperConnectionString(cfg)
	return sqlx.Connect("postgres", dbSource)
}
