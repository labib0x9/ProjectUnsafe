package model

type User struct {
	Role       string `json:""`
	Username   string
	Password   string
	SolvedLabs []string
}

var UserList []User
var Count uint64 = 0
