package shapes

import (
	"math"
)

type form interface {
	area() float64
}

type Rect struct {
	Height float64
	Width  float64
}

func (r Rect) Area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Ray float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Ray, 2)
}
