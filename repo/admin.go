package repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/labib0x9/ProjectUnsafe/model"
)

type AdminRepository interface {
	ListContainer()
	TerminateContainer()
	ListUser() ([]model.User, error)
}

type adminRepo struct {
	db *sqlx.DB
}

func NewAdminRepository(db *sqlx.DB) AdminRepository {
	return &adminRepo{
		db: db,
	}
}

func (r *adminRepo) ListContainer()      {}
func (r *adminRepo) TerminateContainer() {}

func (r *adminRepo) ListUser() ([]model.User, error) {
	var users []model.User
	query := `select * from users`
	if err := r.db.Select(&users, query); err != nil {
		return []model.User{}, err
	}
	return users, nil
}
