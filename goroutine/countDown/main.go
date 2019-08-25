// time.Tick返回一个定时通道。
// 每隔指定的时间就向通道发送一个time.Time的当前时间戳
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Commencing countdown.")
	tick := time.Tick(1 * time.Second)

	for countdown := 10; countdown > 0; countdown-- {
		var t time.Time
		fmt.Printf("%d -- ", countdown)
		t = <-tick
		fmt.Printf("%v\n", t)

	}

	fmt.Println("Lunched!")
}
