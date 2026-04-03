package user

import (
	"net/http"
)

type ProfileResp struct {
	Username string
	Photo    string
}

func (h *Handler) GetProfile(w http.ResponseWriter, r *http.Request) {
	// id, _ := strconv.Atoi(r.PathValue("id"))
	// var profile ProfileResp
	// found := false
	// for _, p := range model.UserList {
	// 	if p.Id == id {
	// 		profile.Username = p.Username
	// 		found = true
	// 		break
	// 	}
	// }

	// if found == false {
	// 	http.Error(w, "User not found", 404)
	// 	return
	// }

	// utils.SendJson(w, profile, http.StatusOK)
}
