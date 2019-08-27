package main

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Create and add some files to the archive.
	var buf bytes.Buffer
	tw := tar.NewWriter(&buf) // Tar file writer, write file to byte.Buffer.

	// Create a struct named files to contain the filename and file content.
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files. "},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling license."},
	}

	// For each file, create the tar file.
	for _, file := range files {

		// Create the file header.
		hdr := &tar.Header{
			Name: file.Name,
			Mode: 0600,
			Size: int64(len(file.Body)),
		}

		// Write header to the file.
		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatal(err)
		}

		// Write the file content.
		if _, err := tw.Write([]byte(file.Body)); err != nil {
			log.Fatal(err)
		}
	}

	// After write all files, close the tar file.
	if err := tw.Close(); err != nil {
		log.Fatal(err)
	}

	// Open and iterate through the files in the archive.
	tr := tar.NewReader(&buf) // Tar file reader. Read tar file from bytes.Buffer.

	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break // End of the tar file.
		}

		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Filename: (%s)\n", hdr.Name)
		fmt.Printf("Format  : (%s)\n", hdr.Format)
		fmt.Printf("Size    : (%d) bytes\n", hdr.Size)
		fmt.Printf("Mode    : (%d)\n", hdr.Mode)
		fmt.Printf("UID     : (%d)\n", hdr.Uid)
		fmt.Printf("GID     : (%d)\n", hdr.Gid)
		fmt.Printf("UNAME   : (%s)\n", hdr.Uname)
		fmt.Printf("GNAME   : (%s)\n", hdr.Gname)
		fmt.Printf("ModTime : (%v)\n", hdr.ModTime)
		fmt.Printf("AccTime : (%v)\n", hdr.AccessTime)
		fmt.Printf("ChaTime : (%v)\n", hdr.ChangeTime)

		fmt.Printf("Contents of %s: \n", hdr.Name)
		if _, err := io.Copy(os.Stdout, tr); err != nil {
			log.Fatal(err)
		}

		fmt.Println()
		fmt.Println()
	}
}
