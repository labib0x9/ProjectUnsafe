package model

import (
	"time"

	"github.com/google/uuid"
)

type Profile struct {
	Id         int64     `json:"id" db:"id"`
	UserId     uuid.UUID `json:"user_id" db:"user_id"`
	ProfilePic string    `json:"profile_pic"   db:"profile_pic"`
	CreatedAt  time.Time `json:"created_at"    db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"    db:"updated_at"`
}
