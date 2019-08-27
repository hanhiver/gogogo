package main

import "fmt"
//import "gogogo/programStructure/importTest2"
import "./mypack"

func main() {
	fmt.Println("OK!")
	hello()
	mypack.Myprint()
}

func hello() {
	fmt.Println("hello in main.")
}

