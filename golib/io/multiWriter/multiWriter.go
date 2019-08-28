package main

import "bytes"
import "strings"
import "io"
import "log"
import "fmt"

func main() {
	r := strings.NewReader("some io.Reader stream to be read. \n")

	var buf1, buf2 bytes.Buffer

	// MultiWriter will write the stream to different buf at the same time.
	w := io.MultiWriter(&buf1, &buf2)

	if _, err := io.Copy(w, r); err != nil {
		log.Fatal(err)
	}

	fmt.Println(buf1.String())
	fmt.Println(buf2.String())
}
