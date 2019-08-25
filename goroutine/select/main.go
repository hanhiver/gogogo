// 在countdown的基础上，让程序可以按回车键终止。
// 为了做到这个，启动一个新的goroutine来监听键盘信号。
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("倒计时开始：")
	tick := time.Tick(1 * time.Second)

	abort := make(chan struct{})
	var t time.Time

	go func() {
		os.Stdin.Read(make([]byte, 1)) //读取单个字符。
		abort <- struct{}{}
	}()

	for countdown := 10; countdown > 0; countdown-- {
		select {
		case t = <-tick:
			fmt.Printf("%d -- %v\n", countdown, t)
		case <-abort:
			fmt.Println("发射终止。")
			return
		}
	}

	fmt.Println("火箭发射！")
}
