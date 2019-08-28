package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

// ReadAtLeast reads from r into buf until it has read at least min bytes.
// It returns the number of bytes copied and an error if fewer bytes were read.
// The error is EOF only if no bytes were read.
func main() {
	r := strings.NewReader("Some io.Reader stream to be read.\n")

	buf := make([]byte, 33)
	if _, err := io.ReadAtLeast(r, buf, 9); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", buf)

	// buffer smaller than minimal read size.
	// err will returns "short buffer"
	shortBuf := make([]byte, 3)
	if _, err := io.ReadAtLeast(r, shortBuf, 5); err != nil {
		fmt.Println("error:", err)
	}

	// minimal read size bigger than io.Reader stream.
	// err will returns "unexpect EOF"
	longBuf := make([]byte, 64)
	if _, err := io.ReadAtLeast(r, longBuf, 64); err != nil {
		fmt.Println("error:", err)
	}
}
