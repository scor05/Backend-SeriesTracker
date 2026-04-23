package backendseriestracker

import (
	"encoding/json"
	"net/http"
	"seriesTracker/src/database"
	"seriesTracker/src/database/models"
)

func postSeries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var serie models.Serie
	err := json.NewDecoder(r.Body).Decode(&serie)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Bad series formatting",
		})
		return
	}

	err = database.Store(serie)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Insertion successful"})

}
