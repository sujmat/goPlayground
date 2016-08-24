package main

import (
	"fmt"
	"math"
)

// Shape interface
type Shape interface {
	area() float64
}

//Square implements Shape
type Square struct {
	side float64
}

func (s Square) area() float64 {
	return s.side * s.side
}

//Circle implements shape
type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Function that operates on a interface
func info(s Shape) {
	fmt.Println(s)
	fmt.Println(s.area())
}

func main() {
	c := Circle{radius: 5}
	s := Square{side: 5}
	info(c)
	info(s)
}
