package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("01234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ\n")

	// SectionReader implements Read, Seek and ReadAt on a section.
	// The example means read source r from 5~17.
	s := io.NewSectionReader(r, 5, 17)

	if _, err := io.Copy(os.Stdout, s); err != nil {
		log.Fatal(err)
	}
	fmt.Println()

	buf := make([]byte, 6)
	// ReadAt will start from index 10.
	if _, err := s.ReadAt(buf, 10); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", buf)

	// Seek will change the index to offset 1 of the start.
	if _, err := s.Seek(1, io.SeekStart); err != nil {
		//if _, err := s.Seek(10, io.SeekStart); err != nil {
		log.Fatal(err)
	}
	if _, err := s.Read(buf); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", buf)
}
