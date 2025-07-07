package handlers

import (
	"encoding/json"
	"net/http"

	"movieactors/db"
	"movieactors/models"
)

func GetNationalities(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// все национальности
	nationalities, err := models.GetAllNationalities(db.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// отправка данных в JSON-формате
	json.NewEncoder(w).Encode(nationalities)
}
