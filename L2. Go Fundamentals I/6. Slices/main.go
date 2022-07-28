package main

import "fmt"

func main() {
	// var fibonacciSlice []int

	// fibonacciSlice[0] = 0
	// fibonacciSlice[1] = 1
	// fibonacciSlice[2] = 1
	// fibonacciSlice[3] = 2
	// fibonacciSlice[4] = 3
	// fibonacciSlice[5] = 5
	// fibonacciSlice[6] = 8
	// fibonacciSlice[7] = 13

	fibonacciSlice := []int{0, 1, 1, 2, 3, 5, 8, 13}

	fmt.Println(fibonacciSlice)
	fmt.Println(fibonacciSlice[0])
	fmt.Println(fibonacciSlice[7])
	fmt.Println(len(fibonacciSlice))
	fmt.Println(fibonacciSlice[0:4])

	fibonacciSlice = append(fibonacciSlice, 21)

	fmt.Println(fibonacciSlice)
	fmt.Println(len(fibonacciSlice))
}
