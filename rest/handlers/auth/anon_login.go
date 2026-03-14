package auth

import (
	"net/http"
	"time"

	"github.com/labib0x9/ProjectUnsafe/model"
	"github.com/labib0x9/ProjectUnsafe/utils"
)

func (h *Handler) AnonLogin(w http.ResponseWriter, r *http.Request) {
	newUser := model.User{
		Role:       "anon",
		Username:   "Guest-",
		Password:   "",
		SolvedLabs: []string{},
		UUID:       utils.Generate_Random_ID(),
		CreatedAt:  time.Now(),
		ExpiredAt:  time.Now().Add(30 * time.Minute),
	}
	newUser.Username += newUser.UUID.String()

	model.UserList = append(model.UserList, newUser)

	data := map[string]any{
		"user": map[string]any{
			"id":          newUser.UUID,
			"username":    newUser.Username,
			"isAdmin":     false,
			"isAnonymous": true,
		},
		"token": "token-blabla",
	}

	utils.SendJson(w, data)
}
