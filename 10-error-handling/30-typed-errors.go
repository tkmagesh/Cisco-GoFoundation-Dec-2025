package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("divide by zero error")

func main() {
	// divisor := 7
	divisor := 0
	q, err := divide(100, divisor)
	if err == nil {
		fmt.Printf("Q : %d\n", q)
	} else if err == ErrDivideByZero {
		fmt.Println("Do not attempt to divide by zero")
	} else {
		fmt.Println("unknown error :", err)
	}
}

func divide(multiplier, divisor int) (quotient int, err error) {
	if divisor == 0 {
		err = ErrDivideByZero
		return
	}
	quotient = multiplier / divisor
	return
}
