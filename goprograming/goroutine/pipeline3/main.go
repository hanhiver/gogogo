package main

import (
	"fmt"
	"time"
)

func counter(out chan<- int) {
	for x := 0; x < 10; x++ {
		out <- x
		time.Sleep(10 * time.Millisecond)
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// counter
	go counter(naturals)

	// squarer
	go squarer(squares, naturals)

	// printer
	for x := range squares {
		fmt.Println(x)
	}

}
