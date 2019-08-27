package popcount

import (
    "fmt"
    "sync"
)

// pc[i]是i的种群统计
var pc [256]byte

// initOnce确保init在并行环境下的执行安全，确保init执行且被执行一次。
var initOnce sync.Once

// func init() {
// 为了确保多线程安全，init()函数必须用initOnce保护。
func initMT() {
    fmt.Println("popcount init.")
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// Popcount返回x的种群统计（置位的个数）
func PopCount(x uint64) int {
    initOnce.Do(initMT)
	return int(pc[byte(x>>(0*8))] +
	           pc[byte(x>>(1*8))] +
	           pc[byte(x>>(2*8))] +
	           pc[byte(x>>(3*8))] +
	           pc[byte(x>>(4*8))] +
	           pc[byte(x>>(5*8))] +
	           pc[byte(x>>(6*8))] +
	           pc[byte(x>>(7*8))])
}

