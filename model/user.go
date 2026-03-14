package model

import (
	"time"

	"github.com/google/uuid"
)

// admin = Admin (ROOT), user = Non-admin, anon = Guest user
type User struct {
	Role       string
	Username   string
	Password   string
	UUID       uuid.UUID
	CreatedAt  time.Time
	ExpiredAt  time.Time
	SolvedLabs []string
}

// // guest user
// type AnonUser struct {
// }

var UserList []User

// var AnonUserList []AnonUser
var Count uint64 = 0
