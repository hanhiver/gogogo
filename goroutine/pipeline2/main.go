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
		for x := 0; x < 10; x++ {
			naturals <- x
			time.Sleep(10 * time.Millisecond)
		}
		close(naturals)
	}()

	// squarer
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// printer
	for x := range squares {
		fmt.Println(x)
	}

}
