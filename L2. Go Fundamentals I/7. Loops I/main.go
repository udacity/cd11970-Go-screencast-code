package main

import "fmt"

func main() {
	// fmt.Println(0)
	// fmt.Println(2)
	// fmt.Println(4)
	// fmt.Println(6)
	// fmt.Println(8)
	// fmt.Println(10)
	// fmt.Println(12)
	// fmt.Println(14)
	// fmt.Println(16)
	// fmt.Println(18)
	// fmt.Println(20)

	for i := 0; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// Alternative
	for i := 0; i <= 20; i += 2 {
		fmt.Println(i)
	}
}
