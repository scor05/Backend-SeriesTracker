package routes

import (
	"encoding/json"
	"net/http"
	"seriesTracker/src/database"
	"seriesTracker/src/database/models"
	"strconv"
)

func PutSeries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid query parameters",
		})
		return
	}

	_, err2 := database.Show(id)
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid series ID",
		})
		return
	}

	var serie models.Serie
	err = json.NewDecoder(r.Body).Decode(&serie)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Bad series formatting",
		})
		return
	}

	err = database.Update(id, serie)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Update successful"})
}
