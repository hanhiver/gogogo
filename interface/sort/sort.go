package main

import (
	"fmt"
	"sort"
)

// Define a type IntSlice to implement the interface of sort.Sort.
type IntSlice []int

// Implementation of the Len(), Less() and Swap() methodes.
func (p IntSlice) Len() int {
	return len(p)
}

func (p IntSlice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p IntSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Call the sort.Sort to sort the int slice.
func main() {
	a := []int{3, 1, 4, 1, 5, 9, 2, 6}
	sort.Sort(IntSlice(a))
	fmt.Printf("%#v\n", a)
}
