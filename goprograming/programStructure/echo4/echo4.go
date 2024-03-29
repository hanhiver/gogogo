// echo4输出命令行参数
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing new line")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}

	x := 1
	y := 2

	x, y = y, x

	fmt.Println(x)
	fmt.Println(y)

}
