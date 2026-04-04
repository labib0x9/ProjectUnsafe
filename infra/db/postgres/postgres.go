package postgres

import "github.com/jmoiron/sqlx"

type PostgreSQL struct {
	*sqlx.DB
}

func New() *PostgreSQL {
	return &PostgreSQL{}
}
