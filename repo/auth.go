package repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/labib0x9/ProjectUnsafe/model"
)

type AuthRepository interface {
	GetByEmail(email string) (model.User, error)
	Create(user model.User) (model.User, error)
}

type authRepo struct {
	dbConn *sqlx.DB
}

func NewAuthRepository(dbConn *sqlx.DB) AuthRepository {
	return &authRepo{
		dbConn: dbConn,
	}
}

func (r *authRepo) GetByEmail(email string) (model.User, error) {
	query := `select * from users where email = $1`
	var user model.User
	if err := r.dbConn.Get(&user, query, email); err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *authRepo) Create(user model.User) (model.User, error) {
	query := `insert into 
		users(username, fullname, email, password_hash, is_verified, role, profile_pic)
		values(:username, :fullname, :email, :password_hash, :is_verified, :role, :profile_pic)
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
