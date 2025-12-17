package main

import "fmt"

// cannot use ":="" in package scope
// x := 100

var unused_var_pkg bool = true

func main() {

	/*
		var userName string
		userName = "Magesh"
	*/

	// Declaration & initialization
	/*
		var userName string = "Magesh"
	*/

	// type inference
	/*
		var userName = "Magesh"
	*/

	// Using := (for declaration, type inference & initialization)
	userName := "Magesh"
	fmt.Printf("Hi %s, Have a nice day!\n", userName)

	// multiple variables
	/*
		var x int
		var y int
		var str string
		var result int
		x = 100
		y = 200
		str = "Sum of %d and %d is %d\n"
		result = x + y
		fmt.Printf(str, x, y, result)
	*/

	// combine the declaration of variables of the same type together
	/*
		var x, y, result int
		var str string
		x = 100
		y = 200
		str = "Sum of %d and %d is %d\n"
		result = x + y
		fmt.Printf(str, x, y, result)
	*/

	// combine the declaration of variables together irrespective of their type
	/*
		var (
			x, y, result int
			str          string
		)
		x = 100
		y = 200
		str = "Sum of %d and %d is %d\n"
		result = x + y
		fmt.Printf(str, x, y, result)
	*/

	// declaration & initialization together
	/*
		var x int = 100
		var y int = 200
		var str string = "Sum of %d and %d is %d\n"
		var result int = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var x, y int = 100, 200
		var str string = "Sum of %d and %d is %d\n"
		var result int = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var (
			x, y   int    = 100, 200
			str    string = "Sum of %d and %d is %d\n"
			result int    = x + y
		)
		fmt.Printf(str, x, y, result)
	*/

	// type inference
	/*
		var (
			x, y   = 100, 200
			str    = "Sum of %d and %d is %d\n"
			result = x + y
		)
		fmt.Printf(str, x, y, result)
	*/

	/*
		var (
			x, y, str = 100, 200, "Sum of %d and %d is %d\n"
			result    = x + y
		)
		fmt.Printf(str, x, y, result)
	*/

	// Using ":="

	x, y, str := 100, 200, "Sum of %d and %d is %d\n"
	result := x + y
	fmt.Printf(str, x, y, result)

	// complex types
	var c1 complex128
	c1 = 5 + 4i
	fmt.Println("c1 = ", c1)
	fmt.Println("real(c1) = ", real(c1))
	fmt.Println("imag(c1) = ", imag(c1))

	// complex type manipulation
	c2 := 6 + 9i
	c3 := c1 + c2
	fmt.Println("c3 = ", c3)

	/*
		var unused_var_fn bool
		unused_var_fn = true
		fmt.Println(unused_var_fn)
	*/
}
