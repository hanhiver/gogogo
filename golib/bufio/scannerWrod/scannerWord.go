package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// An artificial input source.
	const input = "Now is the winter of our discontent,\nMade glorious summer by this sun of York.\n吾等不满之冬，\n已被约克的红日照耀成光荣之夏.\n"

	scanner := bufio.NewScanner(strings.NewReader(input))

	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)

	// Count the words.
	count := 0
	for scanner.Scan() {
		count++
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Reading input: ", err)
	}

	fmt.Printf("\nTotal %d words. \n", count)
}
