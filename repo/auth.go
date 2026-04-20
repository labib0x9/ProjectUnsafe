package repo

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labib0x9/ProjectUnsafe/model"
)

type AuthRepository interface {
	GetByEmail(email string) (model.User, error)
	CreateUser(user model.User) (model.User, error)
	DeleteUserById(id uuid.UUID) error
	DeleteUserEmail(email string) error
	SetVerified(userId uuid.UUID) error
	CreateVerifier(verifier model.Verifier) error
	GetVerifierByHash(tokenHash string) (model.Verifier, error)
	DeleteVerifier(id int64) error
	GetVerifierById(userId uuid.UUID) (model.Verifier, error)
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

func (r *authRepo) CreateUser(user model.User) (model.User, error) {
	query := `insert into 
		users(username, fullname, email, password_hash, is_verified, role, profile_pic, deleted_at)
		values(:username, :fullname, :email, :password_hash, :is_verified, :role, :profile_pic, :deleted_at)
		returning id, username, fullname, email, is_verified, role, created_at
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

func (r *authRepo) DeleteUserById(id uuid.UUID) error {
	query := `delete from users where id = $1`
	_, err := r.dbConn.Exec(query, id)
	return err
}

func (r *authRepo) DeleteUserEmail(email string) error {
	query := `delete from users where email = $1`
	_, err := r.dbConn.Exec(query, email)
	return err
}

func (r *authRepo) CreateVerifier(verifier model.Verifier) error {
	query := `insert into 
		verifier(user_id, token_hash)
		values(:user_id, :token_hash)
	`

	_, err := r.dbConn.NamedExec(query, verifier)
	return err
}

func (r *authRepo) GetVerifierByHash(tokenHash string) (model.Verifier, error) {
	query := `select * from verifier where token_hash = $1 and expire_at > now()`
	var verifier model.Verifier
	if err := r.dbConn.Get(&verifier, query, tokenHash); err != nil {
		return model.Verifier{}, err
	}
	return verifier, nil
}

func (r *authRepo) GetVerifierById(userId uuid.UUID) (model.Verifier, error) {
	query := `select * from verifier where user_id = $1`
	var verifier model.Verifier
	if err := r.dbConn.Get(&verifier, query, userId); err != nil {
		return model.Verifier{}, err
	}
	return verifier, nil
}

func (r *authRepo) SetVerified(userId uuid.UUID) error {
	query := `update users set is_verified = true where id = $1`
	_, err := r.dbConn.Exec(query, userId)
	return err
}

func (r *authRepo) DeleteVerifier(id int64) error {
	query := `delete from verifier where id = $1`
	_, err := r.dbConn.Exec(query, id)
	return err
}
