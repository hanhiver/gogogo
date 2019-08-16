package main

import "fmt"
import "./mypack"

func main() {
	fmt.Println("OK!")
	hello()
}

func hello() {
	fmt.Println("hello in main.")
}

