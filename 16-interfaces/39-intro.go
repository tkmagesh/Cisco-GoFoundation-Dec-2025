package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectange struct {
	Height float64
	Width  float64
}

func (r Rectange) Area() float64 {
	return r.Height * r.Width
}

/*
func PrintArea(x interface{}) {
	switch val := x.(type) {
	case Circle:
		fmt.Println("Area :", val.Area())
	case Rectange:
		fmt.Println("Area :", val.Area())
	default:
		fmt.Println("Invalid type!")
	}
}
*/

/*
func PrintArea(x interface{}) {
	switch val := x.(type) {
	case interface{ Area() float64 }:
		fmt.Println("Area :", val.Area())
	default:
		fmt.Println("Invalid type!")
	}
}
*/

func PrintArea(x interface{ Area() float64 }) {
	fmt.Println("Area :", x.Area())
}

// ver 2.0
type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

/*
Introduce Perimeter() method to all the shape types
	Cirle = 2 * pi * r
	Rectangle = 2 * (h + w)
	Square = 4 * s

Write PrintPerimeter() that prints the perimeter of the given shape
*/

func main() {
	c := Circle{Radius: 12}
	// fmt.Println("Area :", c.Area())
	PrintArea(c)
	PrintPerimeter(c)

	r := Rectange{Height: 12, Width: 14}
	// fmt.Println("Area :", r.Area())
	PrintArea(r)
	PrintPerimeter(r)

	s := Square{Side: 16}
	PrintArea(s)
	PrintPerimeter(s)

}
