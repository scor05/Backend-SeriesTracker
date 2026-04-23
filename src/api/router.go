package main

import (
	"fmt"
	"net/http"
	"seriesTracker/src/api/routes"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:42067")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /series", routes.GetSeries)
	mux.HandleFunc("GET /series/{id}", routes.GetSeriesID)
	mux.HandleFunc("POST /series", routes.PostSeries)
	mux.HandleFunc("PUT /series/{id}", routes.PutSeries)
	mux.HandleFunc("DELETE /series/{id}", routes.Delete)

	fmt.Println("Server listening on port 42069")
	err := http.ListenAndServe(":42069", enableCORS(mux))
	if err != nil {
		fmt.Println("Error listening to port:", err)
	}
}
