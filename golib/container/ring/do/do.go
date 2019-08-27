package main

import (
	"container/ring"
	"fmt"
)

func printElem(p interface{}) {
	fmt.Printf("%d ", p.(int))
}

func main() {
	// Create a new ring of size 5
	r := ring.New(5)

	// Get the length of the ring.
	n := r.Len()

	// Initialize the ring with some integer values.
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	// Iterate through the ring and print its contents.
	r.Do(printElem)
	fmt.Println()

	// Create second ring which will link to the first one.
	r2 := ring.New(3)
	n = r2.Len()
	for j := 0; j < n; j++ {
		r2.Value = 8
		r2 = r2.Next()
	}

	r2.Do(printElem)
	fmt.Println()

	// Link the r and r2 together.
	rs := r.Link(r2)
	rs.Do(printElem)
	fmt.Println()
}
