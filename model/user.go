package model

import (
	"time"
)

// admin = Admin (ROOT), user = Non-admin, anon = Guest user
type User struct {
	Id           int64     `json:"id" db:"id"`
	Username     string    `json:"username"      db:"username"`
	Fullname     string    `json:"fullname"      db:"fullname"`
	Email        string    `json:"email"         db:"email"`
	PasswordHash string    `json:"-"             db:"password_hash"`
	IsVerified   bool      `json:"is_verified"   db:"is_verified"`
	Role         string    `json:"role"          db:"role"`
	ProfilePic   string    `json:"profile_pic"   db:"profile_pic"`
	CreatedAt    time.Time `json:"created_at"    db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"    db:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"    db:"deleted_at"`
}
