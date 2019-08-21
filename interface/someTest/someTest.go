package main

import (
	"io"
	"os"
)

func main() {
	// Test the interface.
	var w io.Writer
	w = os.Stdout
	w.Write([]byte("hello\n"))
}
