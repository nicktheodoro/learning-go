package main

import (
	"fmt"
	"math"
)

func main() {
	r := rectangle{10, 20}
	area(r)
	c := circle{10}
	area(c)
}

type rectangle struct {
	height float64
	width  float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * (c.radius * c.radius)
}

type shape interface {
	area() float64
}

func area(s shape) {
	fmt.Printf("The area of my shape is %0.2f\n", s.area())
}
