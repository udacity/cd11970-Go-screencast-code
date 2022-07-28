package main

import "fmt"

type Baseball struct {
	mass         int
	acceleration int
}

type Football struct {
}

func (b Baseball) getForce() int {
	return b.mass * b.acceleration
}

func (f Football) getForce() int {
	return 50
}

type Force interface {
	getForce() int
}

func compareForce(a, b Force) bool {
	return a.getForce() > b.getForce()
}

func main() {
	b := Baseball{2, 20}
	f := Football{}

	fmt.Println(compareForce(b, f))
}
