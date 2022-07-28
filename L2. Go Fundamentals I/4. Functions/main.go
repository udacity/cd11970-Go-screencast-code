package main

import (
	"fmt"
	"strings"
)

func add(n1, n2 int) int {
	return n1 + n2
}

func sayHello(name string) string {
	return "Hello " + name
}

func sayLoudly(phrase string) string {
	return strings.ToUpper(phrase)
}

func main() {
	fmt.Println(add(2, 8))
	fmt.Println(sayHello("Andrew"))
	fmt.Println(sayLoudly("We are Gophers"))
}
