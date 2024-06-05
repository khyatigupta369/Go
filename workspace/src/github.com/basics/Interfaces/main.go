package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

// value receiver
func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{0, 0, 5}
	rec := Rectangle{10, 5}

	fmt.Printf("Circle: %f\n", getArea(circle))
	fmt.Printf("Rectangle: %f\n", getArea(rec))
}
