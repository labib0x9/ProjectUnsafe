package repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/labib0x9/ProjectUnsafe/model"
)

type UserRepository interface {
	GetProfile(id int) (model.User, error)
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) GetProfile(id int) (model.User, error) {
	query := `select * from users where id = $1`
	var user model.User
	if err := r.db.Get(&user, query, id); err != nil {
		return model.User{}, err
	}
	return user, nil
}