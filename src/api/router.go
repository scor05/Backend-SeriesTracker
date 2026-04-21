package backendseriestracker

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("GET /series", getSeries)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error listening to port:", err)
	}
}
