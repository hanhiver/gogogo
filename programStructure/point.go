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
	fmt.Println(f() == f())

	m := 1
	fmt.Println(incr(&m))
	fmt.Println(m)
}

func f() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p++ // increase the value that pass in to the func.
	return *p
}
