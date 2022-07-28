package main

import (
	"fmt"
	"net/http"
)

func main() {
	// "FileServer returns a handler that serves HTTP requests with the contents of the file system rooted at root"
	// See: https://pkg.go.dev/net/http#FileServer
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	fmt.Println("Server is starting...")
	http.ListenAndServe(":3000", nil)
}
