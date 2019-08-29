package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	content, err := ioutil.ReadFile("readFile.go")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File contents: \n%s\n", content)
}
