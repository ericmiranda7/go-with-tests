package shapes

import (
	"math"
)

type Rectangle struct {
	width  float64
	height float64
}

type Triangle struct {
	base   float64
	height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}

type Shape interface {
	Area() float64
}

func (r Rectangle) Area() float64 {
	return 2 * (r.width + r.height)
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.width + r.height)
}
