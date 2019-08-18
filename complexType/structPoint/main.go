package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

func main() {
	var fullhd Point
	fullhd.x = 1920
	fullhd.y = 1080
	var fullhdp = &fullhd

	fmt.Printf("Width of the fullhd is %d\n", fullhd.x)
	fmt.Printf("Height of the fullhd is %d\n", fullhdp.y)
}
