package main

import "fmt"

func main() {
	var x int
	x = 100

	// Value -> Address
	// Address (pointer) of x
	var xPtr *int
	xPtr = &x
	fmt.Println(xPtr)

	// Address -> Value (dereferencing)
	var x_val int
	x_val = *xPtr // Value at xPtr
	fmt.Println(x_val)

	fmt.Println(x == *(&x))
}
