package main

import (
	"fmt"
	"math"
)

type form interface {
	area() float64
}

type rect struct {
	height float64
	width  float64
}

func (r rect) area() float64 {
	return r.height * r.width
}

type circle struct {
	ray float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.ray, 2)
}

func writeArea(f form) {
	fmt.Printf("The form area is %0.2f\n", f.area())
}

func main() {
	r := rect{10, 15}
	writeArea(r)

	c := circle{10}
	writeArea(c)
}
