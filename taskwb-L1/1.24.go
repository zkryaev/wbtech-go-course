package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func New(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func Distance(p1, p2 *Point) float64 {
	return math.Sqrt(((p2.x - p1.x) * (p2.x - p1.x)) + ((p2.y - p1.y) * (p2.y - p1.y)))
}

func main() {
	a := New(3, 5)
	b := New(7, 2)
	fmt.Println(Distance(a, b))
	fmt.Println(math.Sqrt(61))
}
