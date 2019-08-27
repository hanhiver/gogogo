package main

import (
	"flag"
	"fmt"
	"strings"

	"gogogo/basicTypes/stringLastIndex/basename"
)

var sep = flag.String("s", "/", "separator")

func main() {
	flag.Parse()

	str := basename.Basename(strings.Join(flag.Args(), *sep))
	fmt.Println(str)

}
