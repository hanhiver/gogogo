package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	r1 := strings.NewReader("first reader. \n")
	r2 := strings.NewReader("second reader. \n")
	r3 := strings.NewReader("third reader. \n")

	// MultiReader can read from multiple readers.
	r := io.MultiReader(r1, r2, r3)

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
}
