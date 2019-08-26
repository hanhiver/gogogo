// 构建两个chan和两个goroutine，互相转发一个uint64的值，看看最后的速度多少Mbps
package main

import (
	"fmt"
	"sync"
	"time"
)

func transmit(recv chan uint64, send chan uint64, stop chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("transmit in")

	for {
		select {
		case msg := <-recv:
			//fmt.Printf("Got %d", msg)
			msg++
			send <- msg
		case <-stop:
			count := <-recv
			fmt.Printf("goroutine stopped, %f Mbps in a seconds. \n", float64(count)*2*8/1e6)
			return
		}
	}
}

func main() {
	ch1 := make(chan uint64)
	ch2 := make(chan uint64)
	stop := make(chan struct{})
	tick := time.Tick(4 * time.Second)

	var wg sync.WaitGroup

	go func() {
		wg.Add(1)
		defer wg.Done()

		fmt.Println("Starter in")
		ch1 <- 1
		fmt.Println("Send the starter 1.")
		<-tick
		fmt.Println("Time to stop")
		stop <- struct{}{}
		fmt.Println("Stop signal sent.")
	}()

	wg.Add(1)
	go transmit(ch1, ch2, stop, &wg)
	go transmit(ch2, ch1, stop, &wg)

	wg.Wait()
}
