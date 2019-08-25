// 使用token限制goroutine最大并发数目。
package main

import (
	"fmt"
	"sync"
	"time"
)

func parallelJob(sleepTimes []int) int {
	sizes := make(chan int)

	var wg sync.WaitGroup
	var tokens = make(chan struct{}, 3)

	for _, f := range sleepTimes {
		wg.Add(1)

		go func(s int) {
			defer wg.Done()

			tokens <- struct{}{} // 获取令牌
			fmt.Printf("Sleep for %d seconds.\n", s)

			for i := 0; i < s; i++ {
				time.Sleep(1000 * time.Millisecond)
			}

			fmt.Printf("Sleep for %d seconds Done\n", s)
			sizes <- s
			<-tokens // 释放令牌
		}(f)
	}

	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int
	for size := range sizes {
		total += size
	}

	return total
}

func main() {
	sleepTimes := []int{1, 2, 1, 2, 3, 5, 2, 1, 2, 3}
	parallelJob(sleepTimes)
}
