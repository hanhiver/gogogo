package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// An artificial input source.
	const input = "1234 567890 一百二十 1234567901234567890"

	scanner := bufio.NewScanner(strings.NewReader(input))

	// Create a custom split function by wrapping the existing ScanWords function.
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanWords(data, atEOF)

		fmt.Printf("advance: %v\n", advance)
		for _, c := range token {
			fmt.Printf("%c\t", c)
		}
		fmt.Println()

		if err == nil && token != nil {
			// Convert the token to 10 based 32 bits int.
			_, err = strconv.ParseInt(string(token), 10, 32)
		}

		return
	}

	// Set the split function for the scanning operation.
	scanner.Split(split)

	// Validate the input.
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s\n", err)
	}
}
