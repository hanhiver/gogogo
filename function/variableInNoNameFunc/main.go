package main

import (
	"fmt"
)

func main() {
	// 准备一个列表，我们将打印列表中的数字并且将其全部翻倍再次输出。
	myList := []int{3, 1, 4, 1, 5, 9, 2, 6}

	// 定义包含函数的slice进行处理。
	var doubleFuncs []func()

	for _, item := range myList {

		//此处一定要新建一个temp来暂存信息，item会在后续处理中被不断更新。
		temp := item

		fmt.Printf("列出值： %d\n", item)
		doubleFuncs = append(doubleFuncs, func() {
			fmt.Printf("翻倍列出值(TEMP): %d\n", 2*temp)

			// 此处错误，如果用item会被不断更新，最后只有一个值。
			fmt.Printf("翻倍列出值(ITEM): %d\n", 2*item)
		})
	}

	for _, doubleItem := range doubleFuncs {
		doubleItem()
	}

}
