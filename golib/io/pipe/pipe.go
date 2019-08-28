package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	// Create a pair of pipe one for read and one for write.
	r, w := io.Pipe()

	go func() {
		fmt.Fprint(w, "some text to be read. \n")
		w.Close()
	}()

	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	fmt.Print(buf.String())

}
