package repo

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labib0x9/ProjectUnsafe/infra/cache/redis"
	"github.com/labib0x9/ProjectUnsafe/model"
)

var ctx = context.Background()

type AuthRepository interface {
	GetByEmail(email string) (model.User, error)
	GetUserById(id uuid.UUID) (model.User, error)
	CreateUser(user model.User) (model.User, error)
	DeleteUserById(id uuid.UUID) error
	DeleteUserEmail(email string) error
	SetVerified(userId uuid.UUID) error
	CreateVerifier(verifier model.Verifier) error
	GetVerifierByHash(tokenHash string) (model.Verifier, error)
	DeleteVerifier(id int64) error
	GetVerifierById(userId uuid.UUID) (model.Verifier, error)
	Set(key string, value string, expire time.Duration) error
	Get(key string) (string, error)
	GetResetToken(id uuid.UUID) (model.Reseter, error)
	UpdateResetToken(reseter model.Reseter) error
	CreateResetToken(reseter model.Reseter) error
	GetResetTokenByToken(token string) (model.Reseter, error)
	UpdateUserPassword(id uuid.UUID, passHash string) error
	DeleteTokenById(id int64) error
}

type authRepo struct {
	dbConn *sqlx.DB
	cache  *redis.Redis
}

func NewAuthRepository(
	dbConn *sqlx.DB,
	cache *redis.Redis,
) AuthRepository {
	return &authRepo{
		dbConn: dbConn,
		cache:  cache,
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

func (r *authRepo) GetUserById(id uuid.UUID) (model.User, error) {
	query := `select * from users where id = $1`
	var user model.User
	if err := r.dbConn.Get(&user, query, id); err != nil {
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

func (r *authRepo) Set(key string, value string, expire time.Duration) error {
	return r.cache.Client.Set(
		ctx,
		key,
		value,
		expire,
	).Err()
}

func (r *authRepo) Get(key string) (string, error) {
	return r.cache.Client.Get(ctx, key).Result()
}

func (r *authRepo) GetResetToken(id uuid.UUID) (model.Reseter, error) {
	query := `select * from reseter where user_id = $1 and expire_at > now()`
	var reseter model.Reseter
	if err := r.dbConn.Get(&reseter, query, id); err != nil {
		return model.Reseter{}, err
	}
	return reseter, nil
}

// no use case
func (r *authRepo) UpdateResetToken(reseter model.Reseter) error {
	return nil
}

func (r *authRepo) CreateResetToken(reseter model.Reseter) error {
	query := `insert into 
		reseter(user_id, token_hash)
		values(:user_id, :token_hash)
	`

	_, err := r.dbConn.NamedExec(query, reseter)
	return err
}

func (r *authRepo) GetResetTokenByToken(token string) (model.Reseter, error) {
	query := `select * from reseter where token_hash = $1 and expire_at > now()`
	var reseter model.Reseter
	if err := r.dbConn.Get(&reseter, query, token); err != nil {
		return model.Reseter{}, err
	}
	return reseter, nil
}

func (r *authRepo) DeleteTokenById(id int64) error {
	query := `delete from reseter where id = $1`
	_, err := r.dbConn.Exec(query, id)
	return err
}

func (r *authRepo) UpdateUserPassword(id uuid.UUID, passHash string) error {
	query := `update users set password_hash = $1 where id = $2`
	_, err := r.dbConn.Exec(query, passHash, id)
	return err
}
