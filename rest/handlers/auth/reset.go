package auth

import (
	"net/http"
)

type ResetResp struct {
	Email string
}

func (h *Handler) Reset(w http.ResponseWriter, r *http.Request) {
	// var resp ResetResp
	// decoder := json.NewDecoder(r.Body)
	// if err := decoder.Decode(&resp); err != nil {
	// 	http.Error(w, "Bad request", http.StatusBadRequest)
	// 	return
	// }

	// if err := h.authRepo.Reset(resp.Email); err != nil {
	// 	http.Error(w, "internal server error", http.StatusInternalServerError)
	// 	return
	// }

	// utils.SendJson(w, "reset url", http.StatusOK)

}
