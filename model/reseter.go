package model

import (
	"time"

	"github.com/google/uuid"
)

type Reseter struct {
	Id        int64     `json:"id" db:"id"`
	UserId    uuid.UUID `json:"user_id" db:"user_id"`
	Token     string    `json:"token_hash" db:"token_hash"`
	Used      bool      `json:"used" db:"used"`
	CreatedAt time.Time `json:"created_at"    db:"created_at"`
	ExpireAt  time.Time `json:"expire_at"    db:"expire_at"`
}
