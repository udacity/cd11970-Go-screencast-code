package main

import "fmt"

func main() {
	// var fibonacciArray [8]int

	// fibonacciArray[0] = 0
	// fibonacciArray[1] = 1
	// fibonacciArray[2] = 1
	// fibonacciArray[3] = 2
	// fibonacciArray[4] = 3
	// fibonacciArray[5] = 5
	// fibonacciArray[6] = 8
	// fibonacciArray[7] = 13

	fibonacciArray := [8]int{0, 1, 1, 2, 3, 5, 8, 13}

	fmt.Println(fibonacciArray)
	fmt.Println(fibonacciArray[0])
	fmt.Println(fibonacciArray[7])
	fmt.Println(len(fibonacciArray))
	fmt.Println(fibonacciArray[0:4])
}
