package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}

	fmt.Println("Test Write()")
	w := csv.NewWriter(os.Stdout)

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatal("Error writing record to csv: ", err)
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
	fmt.Println()

	fmt.Println("Test WriteAll()")
	w.WriteAll(records) // call Flush internally.

	if err := w.Error(); err != nil {
		log.Fatal("Error writing csv:", err)
	}
}
