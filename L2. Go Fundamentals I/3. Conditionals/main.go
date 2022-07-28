package main

import "fmt"

func main() {
	language := "Go"

	if language == "Java" {
		fmt.Println("language is Java")
	} else if language == "C" {
		fmt.Println("language is C")
	} else if language == "C++" {
		fmt.Println("language is C++")
	} else {
		fmt.Println("language is not Java, C, or C++")
	}
}
