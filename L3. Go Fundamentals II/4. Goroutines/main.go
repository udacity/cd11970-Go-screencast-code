package main

import (
	"fmt"
	"time"
)

func showMessage(message string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(message)
	}
}

func main() {
	go showMessage("Go is a great programming language.")

	showMessage("Welcome, Gophers!")
}
