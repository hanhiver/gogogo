package main

import (
	"flag"
	"fmt"
	"time"
)

var p = flag.Duration("p", 1*time.Second, "Sleep Period")

func main() {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *p)
	time.Sleep(*p)
	fmt.Println()
}
