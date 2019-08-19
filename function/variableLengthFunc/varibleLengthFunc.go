// 测试可变长度函数。
package main

import (
	"fmt"
)

func main() {

	fmt.Println(sum())
	fmt.Println(sum(3))
	fmt.Println(sum(1, 2, 3, 4, 5))

	values := []int{1, 2, 3, 4, 5}
	// 如果所有的参数已经放到一个列表中，调用的时候添加...
	fmt.Println(sum(values...))

	// 变长函数的类型和带有普通slice参数的函数不同。
	fmt.Printf("typeof(ff) = %#v \ntypeof(gg) = %#v \n", ff, gg)
}

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}

	return total
}

func ff(...int) {}
func gg([]int)  {}
