package main

import (
	"fmt"
)

func main() {
	// var x interface{}
	var x any

	// x = 100
	// x = 99.99
	// x = true
	// x = "Elit non laborum non dolor dolor qui sint enim ea laboris."
	// x = struct{}{}

	// x = 100
	x = "Ut consequat elit eiusmod ea cillum occaecat esse id sint."

	// compilation error
	// y := x * 2

	// MAY result in runtime panic if x does NOT have an int value
	// y := x.(int) * 2

	// To prevent the possibilities of runtime panic
	// type assertion using "if"
	if val, ok := x.(int); ok {
		y := val * 2
		fmt.Println(y)
	} else {
		fmt.Println("x is NOT an int")
	}

	// type assertion using "type-switch"
	// x = 100
	// x = 99.99
	// x = true
	// x = "Elit non laborum non dolor dolor qui sint enim ea laboris."
	x = struct{}{}
	switch val := x.(type) {
	case int:
		fmt.Println("x is an int, x * 2 =", val*2)
	case string:
		fmt.Println("x is a string, len(x) =", len(val))
	case bool:
		fmt.Println("x is a boolean, !x =", !val)
	case float64:
		fmt.Println("x is a float64, x * 0.977 =", val*0.977)
	default:
		fmt.Println("x is of unknown type, x =", x)
	}

}
