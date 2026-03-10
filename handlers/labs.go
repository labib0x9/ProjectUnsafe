package handlers

import (
	"log/slog"
	"net/http"

	"github.com/labib0x9/ProjectUnsafe/model"
	"github.com/labib0x9/ProjectUnsafe/utils"
)

func GetAllLabs(w http.ResponseWriter, r *http.Request) {
	utils.SendJson(w, model.LabList)
}

func GetLabByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var lab model.Lab
	found := false
	for _, tempLab := range model.LabList {
		if tempLab.Id == id {
			lab = tempLab
			found = true
			break
		}
	}

	if found == false {
		w.WriteHeader(404)
		return
	}
	utils.SendJson(w, lab)
}

func StartLab(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, 256)
	r.Body.Read(body)
	slog.Info(string(body))
}

func ResetLab(w http.ResponseWriter, r *http.Request)     {}
func TerminateLab(w http.ResponseWriter, r *http.Request) {}
