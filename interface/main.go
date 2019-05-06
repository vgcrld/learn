package main

import (
	"fmt"
	"math"
)

// Shape make area available to each
// Basically any struct that implements area
type Shape interface {
	area() float64
}

// Circle struct define center point and the radius
type Circle struct {
	x, y, radius float64
}

// Rectangle struct defines the width and hight
type Rectangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rectangle{width: 10, height: 5}

	fmt.Printf("Circle Area: %f\n", getArea(circle))
	fmt.Printf("Rectangle Area: %f\n", getArea(rectangle))
}
