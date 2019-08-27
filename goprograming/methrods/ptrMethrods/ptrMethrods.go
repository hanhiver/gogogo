package main

import (
	"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
}

func (p Point) Distance(q Point) float64 {
	dX := q.X - p.X
	dY := q.Y - p.Y
	return math.Sqrt(dX*dX + dY*dY)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	r := Point{1, 2}
	rptr := &r
	rptr.ScaleBy(2)
	fmt.Println(*rptr)

	// 虽然ScaleBy是绑定指针类型的方法，但是用变量类型调用，会自动完成取指针的隐式转换。
	r.ScaleBy(2)
	fmt.Println(r)

	scaleByFromR := r.ScaleBy
	scaleByFromR(2)
	fmt.Println(r)

	distance := Point.Distance
	p := Point{2, 5}
	fmt.Println(distance(r, p))

}
