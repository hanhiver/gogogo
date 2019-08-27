package main

import "fmt"
import "gogogo/programStructure/importTest/mypack"

func main() {
	fmt.Println("OK!")
	hello()
	mypack.Myprint()
}

func hello() {
	fmt.Println("hello in main.")
}

