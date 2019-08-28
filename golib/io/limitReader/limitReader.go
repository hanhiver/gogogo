package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read.")

	// LimitReader returns a reader that reads from r but stop with EOF after n bytes.
	lr := io.LimitReader(r, 9)

	if _, err := io.Copy(os.Stdout, lr); err != nil {
		log.Fatal(err)
	}
	fmt.Println()
}
