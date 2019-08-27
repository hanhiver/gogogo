// 使用sync.WaitGroup来同步goroutine的运行。
// WaitGroup.Add(1)添加一个正在运行的进程数。
// WaitGroup.Done()释放一个正在运行的进程数。
package main

import (
	"fmt"
	"sync"
	"time"
)

func parallelJob(sleepTimes []int) int {
	sizes := make(chan int)

	var wg sync.WaitGroup

	for _, f := range sleepTimes {
		wg.Add(1)

		go func(s int) {
			defer wg.Done()

			fmt.Printf("Sleep for %d seconds.\n", s)

			for i := 0; i < s; i++ {
				time.Sleep(1000 * time.Millisecond)
			}

			fmt.Printf("Sleep for %d seconds Done\n", s)
			sizes <- s
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
	sleepTimes := []int{1, 2, 1, 2, 3, 5}
	parallelJob(sleepTimes)
}
