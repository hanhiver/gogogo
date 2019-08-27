package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	in := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
`
	r := csv.NewReader(strings.NewReader(in))

	fmt.Println("Test the Read()")
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record)
	}
	fmt.Println()

	fmt.Println("Test the ReadAll()")
	r2 := csv.NewReader(strings.NewReader(in))
	records, err2 := r2.ReadAll()
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(records)
	fmt.Println()

	in2 := `first_name;last_name;username
"Rob";"Pike";rob
# lines beginning with a # character are ignored
Ken;Thompson;ken
"Robert";"Griesemer";"gri"
`
	fmt.Println("Test the Read() with option")
	r3 := csv.NewReader(strings.NewReader(in2))
	r3.Comma = ';'
	r3.Comment = '#'

	for {
		record3, err3 := r3.Read()
		if err3 == io.EOF {
			break
		}
		if err3 != nil {
			log.Fatal(err3)
		}

		fmt.Println(record3)
	}
	fmt.Println()

}
