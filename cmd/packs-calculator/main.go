package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sarataha/packs-calculator/api"
)

func main() {
	router := mux.NewRouter()

	// Serve static files
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/index.html")
	})

	router.HandleFunc("/api/packs", api.PacksHandler).Methods("GET")

	fmt.Println("Server is running on :8080...")
	http.ListenAndServe(":8080", router)
}
