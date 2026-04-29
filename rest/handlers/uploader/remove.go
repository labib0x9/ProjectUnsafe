package uploader

import "net/http"

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	h.uploaderRepo.Delete()
}
