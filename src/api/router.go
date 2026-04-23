package main

import (
	"fmt"
	"net/http"
	"seriesTracker/src/api/routes"
)

func main() {

	http.HandleFunc("GET /series", routes.GetSeries)
	http.HandleFunc("GET /series/{id}", routes.GetSeriesID)
	http.HandleFunc("POST /series", routes.PostSeries)
	http.HandleFunc("PUT /series/{id}", routes.PutSeries)
	http.HandleFunc("DELETE /series/{id}", routes.Delete)

	fmt.Println("Server listening on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error listening to port:", err)
	}
}
