package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	var buf bytes.Buffer

	// TeeReader returns a Reader that writes to w what it reads from.
	// All reads from r performed through it are matched what corresponding writes to w.
	// There is no internal buffering.
	// The write must complete before the read completes.
	// Any error encountered while writing is reported as a read error.
	tee := io.TeeReader(r, &buf)

	printall := func(r io.Reader) {
		b, err := ioutil.ReadAll(r)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", b)
	}

	printall(tee)
	printall(&buf)
}
