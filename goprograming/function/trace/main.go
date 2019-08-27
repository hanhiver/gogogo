package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	bigSlowOperation()
}

func bigSlowOperation() {
	fmt.Println("bigSlowOperation start.")
	defer trace("bigSlowOperation")()
	fmt.Println("defer set. ")

	time.Sleep(1 * time.Second)
	fmt.Println("bigSlowOperation finished.")
	panic("Something Wrong!")
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}
