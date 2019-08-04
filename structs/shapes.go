package structs

import "math"

// Shape with area
type Shape interface {
	area() float64
}

// Rectangle with width and height
type Rectangle struct {
	width  float64
	height float64
}

func (rectangle Rectangle) area() float64 {
	return rectangle.width * rectangle.height
}

// Circle with radius
type Circle struct {
	radius float64
}

func (circle Circle) area() float64 {
	return math.Pi * math.Pow(circle.radius, 2)
}

// Triangle with base and height
type Triangle struct {
	base   float64
	height float64
}

func (triangle Triangle) area() float64 {
	return triangle.base * triangle.height / 2
}
