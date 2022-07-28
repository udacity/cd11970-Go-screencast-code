package main

import "fmt"

func main() {
	dictionary := map[string]string{
		"Go":     "A programming language created by Google.",
		"Gopher": "A software engineer who builds with Go.",
		"Golang": "Another name for Go.",
	}

	fmt.Println(dictionary)
	fmt.Println(dictionary["Go"])

	// Reassignment
	dictionary["Gopher"] = "The fuzzy mascot for Go."

	// New key-value pair
	dictionary["Map"] = "An unordered data structure with key-value pairs."

	fmt.Println(dictionary)

	for word, definition := range dictionary {
		fmt.Println("The definition of", word, "is:", definition)
	}

	delete(dictionary, "Map")

	fmt.Println(dictionary)
}
