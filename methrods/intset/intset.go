// IntSet描述了一个结构，其中用每一位表示一个非负整数是否在结构中。
// 比如对与0， 那么就将第0位设置为1，对与123就将第123位设置为1。
// 当结构中的位数不够的时候，选择在结构内添加更多的unit64数字。
package main

import (
	"bytes"
	"fmt"
)

// IntSet是一个包含非负整形的集合
// 零值代表为空。
type IntSet struct {
	words []uint64
}

// Has方法返回是否存在非负整形x
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add添加非负数x到集合中去。
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith将会对s和t做并集并将结果存储在s中
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String方式输出。
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')

	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String())
	fmt.Printf("%v\n", x.words)

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String())
	fmt.Printf("%v\n", y.words)

	x.UnionWith(&y)
	fmt.Println(x.String())
	fmt.Printf("%v\n", x.words)

	fmt.Println(x.Has(9), x.Has(123))
}
