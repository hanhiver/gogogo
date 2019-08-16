package main

import (
	"fmt"

	"./popcount"
)

func main() {
	var x uint64 = 186
	res := popcount.PopCount(x)

	fmt.Printf("%d\n", res)
}
