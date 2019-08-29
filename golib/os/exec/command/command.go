package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("tr", "a-z", "A-Z")
	cmd.Stdin = strings.NewReader("some input")

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %q\n", out.String())

	out.Reset()
	cmd = exec.Command("echo", "${SHELL}")
	cmd.Env = append(os.Environ(),
		"FOO=duplicate_value", // Duplicated, ignored.
		"FOO=actual_value",    // This value used.
	)
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Echo output: %s", out.String())
}
