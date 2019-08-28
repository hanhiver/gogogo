package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("Some io.Reader stream to be read. ")

	// CopyN will only read certain number from the source stream.
	if _, err := io.CopyN(os.Stdout, r, 5); err != nil {
		log.Fatal(err)
	}
}
