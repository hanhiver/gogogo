package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {
	content := []byte("Temporary file's content. \n")

	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Tempfile %s created. \n", tmpfile.Name())

	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}

	time.Sleep(10 * time.Second)

	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}

}
