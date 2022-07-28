package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var dictionary = map[string]string{
	"Go":     "A programming language created by Google.",
	"Gopher": "A software engineer who builds with Go.",
	"Golang": "Another name for Go.",
}

func getDictionary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(dictionary)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/dictionary", getDictionary).Methods("GET")

	fmt.Println("Server is starting on port 3000...")
	http.ListenAndServe(":3000", router)
}
