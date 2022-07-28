package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world :]")
}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Contact Info</h1>")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/contact", contact)

	fmt.Println("Server is starting...")
	http.ListenAndServe(":3000", nil)
}
