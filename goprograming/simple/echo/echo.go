package main

import (
	"fmt"
	"os"
	"strings"
)

func echo1() {
	var s, sep string
	for i:=0; i<len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println("echo1 " + s)
}

func echo2() {
	var s, sep string
	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println("echo2 " + s)
}

func echo3() {
	fmt.Println("echo3 " + strings.Join(os.Args[0:], " "))
}

func main() {
	echo1()
	echo2()
	echo3()
}



