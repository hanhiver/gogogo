package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

func main() {
	var fullhd1 Point
	// 直接按照成员初始化
	fullhd1.x = 1920
	fullhd1.y = 1080

	fullhd := Point{x: 1920, y: 1080}
	var fullhdp = &fullhd

	fmt.Printf("Width of the fullhd is %#v\n", fullhd.x)
	fmt.Printf("Height of the fullhd is %#v\n", fullhdp.y)

	fmt.Printf("The structure: %#v\n", fullhd)
}
