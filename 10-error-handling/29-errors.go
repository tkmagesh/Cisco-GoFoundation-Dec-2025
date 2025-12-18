package main

import (
	"errors"
	"fmt"
)

func main() {
	// divisor := 7
	divisor := 0
	q, err := divide(100, divisor)
	if err == nil {
		fmt.Printf("Q : %d\n", q)
	} else {
		fmt.Println("error :", err)
	}
}

/*
func divide(multiplier, divisor int) (int, error) {
	var quotient int // initialized to 0
	var err error // intialized to nil
	if divisor == 0 {
		err = errors.New("divisor cannot be zero")
		return quotient, err
	}
	quotient = multiplier / divisor
	return quotient, err
}
*/

func divide(multiplier, divisor int) (quotient int, err error) {
	if divisor == 0 {
		err = errors.New("divisor cannot be zero")
		return
	}
	quotient = multiplier / divisor
	return
}
