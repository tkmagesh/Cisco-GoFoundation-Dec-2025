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

type AreaFinder interface {
	Area() float64
}

func PrintArea(x AreaFinder) {
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
// ver 3.0
Introduce Perimeter() method to all the shape types
	Cirle = 2 * pi * r
	Rectangle = 2 * (h + w)
	Square = 4 * s

Write PrintPerimeter() that prints the perimeter of the given shape
*/

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (r Rectange) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

func (s Square) Perimeter() float64 {
	return 4 * s.Side
}

type PerimeterFinder interface {
	Perimeter() float64
}

func PrintPerimeter(x PerimeterFinder) {
	fmt.Println("Perimeter :", x.Perimeter())
}

/*
func PrintShapeStats(x interface {
	interface{ Area() float64 }      // for PrintArea()
	interface{ Perimeter() float64 } // for PrintPerimeter()
}) {
	PrintArea(x)      // x should be interface{ Area() float64 }
	PrintPerimeter(x) // x should be interface { Perimeter() float64 }
}
*/

/*
func PrintShapeStats(x interface {
	AreaFinder      // for PrintArea()
	PerimeterFinder // for PrintPerimeter()
}) {
	PrintArea(x)      // x should be interface{ Area() float64 }
	PrintPerimeter(x) // x should be interface { Perimeter() float64 }
}
*/

type ShapeStatsFinder interface {
	AreaFinder      // for PrintArea()
	PerimeterFinder // for PrintPerimeter()
}

func PrintShapeStats(x ShapeStatsFinder) {
	PrintArea(x)      // x should be interface{ Area() float64 }
	PrintPerimeter(x) // x should be interface { Perimeter() float64 }
}

/*
func PrintShapeStats(x interface {
	Area() float64      // for PrintArea()
	Perimeter() float64 // for PrintPerimeter()
}) {
	PrintArea(x)      // x should be interface{ Area() float64 }
	PrintPerimeter(x) // x should be interface { Perimeter() float64 }
}
*/

func main() {
	c := Circle{Radius: 12}
	// fmt.Println("Area :", c.Area())
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/
	PrintShapeStats(c)

	r := Rectange{Height: 12, Width: 14}
	// fmt.Println("Area :", r.Area())
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/
	PrintShapeStats(r)

	s := Square{Side: 16}
	/*
		PrintArea(s)
		PrintPerimeter(s)
	*/
	PrintShapeStats(s)
}
