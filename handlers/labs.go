package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/labib0x9/ProjectUnsafe/model"
)

func GetAllLabs(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.Encode(model.LabList)
}
