package main

import (
	"fmt"
)

func main() {
	x, y := 1, 2
	fmt.Println(add(x, y))

}

func add(a, b int) int {
	return a + b
}

