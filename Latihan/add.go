package main

import (
	"fmt"
)

func main() {
	x := 6
	y := 5
	add(x, y)
	min(x, y)
	multi(x, y)
	div(x, y)
}

func add(x int, y int) {
	fmt.Println(x + y)
}

func min(x int, y int) {
	fmt.Println(x - y)
}

func multi(x int, y int) {
	fmt.Println(x * y)
}

func div(x int, y int) {
	fmt.Println(x / y)
}
