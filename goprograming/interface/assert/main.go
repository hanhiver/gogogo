package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout
	f := w.(*os.File)
	fmt.Println(f)

	// This will crash.
	//c := w.(*bytes.Buffer)
	f, ok := w.(*os.File)
	fmt.Println(f, ok)
	b, ok := w.(*bytes.Buffer)
	fmt.Println(b, ok)
}
