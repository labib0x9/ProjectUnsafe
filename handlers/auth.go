package handlers

import (
	"net/http"
	"strconv"

	"github.com/labib0x9/ProjectUnsafe/model"
)

func AnonLogin(w http.ResponseWriter, r *http.Request) {
	Id := strconv.FormatUint(model.Count+1, 10)
	model.Count += 1
	newUser := model.User{
		Role:       "anon",
		Username:   "Anon" + Id,
		Password:   "",
		SolvedLabs: []string{},
	}

	model.UserList = append(model.UserList, newUser)
}
