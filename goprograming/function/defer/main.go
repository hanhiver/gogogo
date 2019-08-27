package main

import (
	"fmt"
)

func main() {
	fmt.Println("Test with true")
	deferTest(true)

	fmt.Println("\nTest with false")
	deferTest(false)

}

func deferTest(isok bool) {
	fmt.Println("In the deferTest() func. ")

	if isok {
		fmt.Println("I am not OK, I will quit.")
		return
	}
	defer func() {
		fmt.Println("TRUE Don't forgot to run defer. ")
	}()

	if !isok {
		fmt.Println("I am fine, quit normally. ")
		return
	}
	defer func() {
		fmt.Println("FALSE Don't forgot to run defer. ")
	}()

}
