package repo

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labib0x9/ProjectUnsafe/model"
)

type ReseterRepo interface {
	GetById(id uuid.UUID) (model.Reseter, error)
	Update(reseter model.Reseter) error
	Create(reseter model.Reseter) error
	GetByToken(token string) (model.Reseter, error)
	DeleteById(id int64) error
}

type reseterRepo struct {
	dbConn *sqlx.DB
}

func NewReseterRepo(
	dbConn *sqlx.DB,
) ReseterRepo {
	return &reseterRepo{
		dbConn: dbConn,
	}
}

func (r *reseterRepo) GetById(id uuid.UUID) (model.Reseter, error) {
	query := `select * from reseter where user_id = $1 and expire_at > now()`
	var reseter model.Reseter
	if err := r.dbConn.Get(&reseter, query, id); err != nil {
		return model.Reseter{}, err
	}
	return reseter, nil
}

// no use case
func (r *reseterRepo) Update(reseter model.Reseter) error {
	return nil
}

func (r *reseterRepo) Create(reseter model.Reseter) error {
	query := `insert into 
		reseter(user_id, token_hash)
		values(:user_id, :token_hash)
	`

	_, err := r.dbConn.NamedExec(query, reseter)
	return err
}

func (r *reseterRepo) GetByToken(token string) (model.Reseter, error) {
	query := `select * from reseter where token_hash = $1 and expire_at > now()`
	var reseter model.Reseter
	if err := r.dbConn.Get(&reseter, query, token); err != nil {
		return model.Reseter{}, err
	}
	return reseter, nil
}

func (r *reseterRepo) DeleteById(id int64) error {
	query := `delete from reseter where id = $1`
	_, err := r.dbConn.Exec(query, id)
	return err
}
