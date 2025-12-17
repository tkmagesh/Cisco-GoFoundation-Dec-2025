package main

import "fmt"

func main() {
	multiplier, divisor := 100, 7
	/*
		q, r := divide(multiplier, divisor)
		fmt.Printf("Dividing %d by %d, quotient = %d and remainder = %d\n", multiplier, divisor, q, r)
	*/

	// use "_" as a placeholder to receive a value that we don't want to use
	q, _ := divide(multiplier, divisor)
	fmt.Printf("Dividing %d by %d, quotient = %d \n", multiplier, divisor, q)
}

/*
func divide(multiplier, divisor int) (int, int) {
	quotient := multiplier / divisor
	remainder := multiplier % divisor
	return quotient, remainder
}
*/

// named results
func divide(multiplier, divisor int) (quotient int, remainder int) { // quotient & remainder will by default be initialized with the 'zero' value of 'int' type
	quotient = multiplier / divisor
	remainder = multiplier % divisor
	return
}
