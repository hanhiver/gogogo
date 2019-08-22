package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// counter
	go func() {
		for x := 0; ; x++ {
			naturals <- x
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// squarer
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	// printer
	for {
		fmt.Println(<-squares)
	}

}
