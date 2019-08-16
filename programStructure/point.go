package main

import (
	"fmt"
)

func main() {
	x := 1
	p := &x
	fmt.Println(*p)
	*p = 2
	fmt.Println(x)

	pv := f()
	fmt.Println(*pv)
}

func f() *int {
	v := 1
	return &v
}
