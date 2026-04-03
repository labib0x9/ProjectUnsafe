package repo

import "github.com/jmoiron/sqlx"

type UserRepository interface {
	AnonLogin()
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) userRepo {
	return userRepo{db: db}
}
