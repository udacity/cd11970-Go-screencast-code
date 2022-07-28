package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	http.HandleFunc("/", getDictionary)

	fmt.Println("Server is starting on port 3000...")
	http.ListenAndServe(":3000", nil)
}
