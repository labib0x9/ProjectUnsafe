package repo

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labib0x9/ProjectUnsafe/model"
)

type VerifierRepo interface {
	Create(verifier model.Verifier) error
	GetByHash(tokenHash string) (model.Verifier, error)
	Delete(id int64) error
	GetById(userId uuid.UUID) (model.Verifier, error)
}

type verifierRepo struct {
	dbConn *sqlx.DB
}

func NewVerifierRepo(
	dbConn *sqlx.DB,
) VerifierRepo {
	return &verifierRepo{
		dbConn: dbConn,
	}
}

func (r *verifierRepo) Create(verifier model.Verifier) error {
	query := `insert into 
		verifier(user_id, token_hash)
		values(:user_id, :token_hash)
	`

	_, err := r.dbConn.NamedExec(query, verifier)
	return err
}

func (r *verifierRepo) GetByHash(tokenHash string) (model.Verifier, error) {
	query := `select * from verifier where token_hash = $1 and expire_at > now()`
	var verifier model.Verifier
	if err := r.dbConn.Get(&verifier, query, tokenHash); err != nil {
		return model.Verifier{}, err
	}
	return verifier, nil
}

func (r *verifierRepo) GetById(userId uuid.UUID) (model.Verifier, error) {
	query := `select * from verifier where user_id = $1`
	var verifier model.Verifier
	if err := r.dbConn.Get(&verifier, query, userId); err != nil {
		return model.Verifier{}, err
	}
	return verifier, nil
}

func (r *verifierRepo) Delete(id int64) error {
	query := `delete from verifier where id = $1`
	_, err := r.dbConn.Exec(query, id)
	return err
}
