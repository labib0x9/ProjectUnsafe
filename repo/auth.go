package repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/labib0x9/ProjectUnsafe/model"
)

type AuthRepository interface {
	AnonLogin(user model.User) (model.User, error)
}

type authRepo struct {
	dbConn *sqlx.DB
}

func NewAuthRepository(dbConn *sqlx.DB) AuthRepository {
	return &authRepo{
		dbConn: dbConn,
	}
}

func (r *authRepo) AnonLogin(user model.User) (model.User, error) {
	query := `
        insert into users (username, email, role, is_verified, deleted_at)
        values (:username, :email, :role, :is_verified, :deleted_at)
        returning *
    `

	rows, err := r.dbConn.NamedQuery(query, user)
	if err != nil {
		return model.User{}, err
	}
	defer rows.Close()

	var created model.User
	if rows.Next() {
		if err := rows.StructScan(&created); err != nil {
			return model.User{}, err
		}
	}
	return created, nil
}
