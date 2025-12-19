/*
Create a custom error type (replacing ErrDivideByZero) that has the following information
	multiplier
	divisor
	error message as a string
*/

package main

import (
	"fmt"
)

// var ErrDivideByZero error = errors.New("divide by zero error")
type ErrDivideByZero struct {
	Multiplier int
	Divisor    int
}

func NewErrDivideByZero(multiplier, divisor int) *ErrDivideByZero {
	return &ErrDivideByZero{
		Multiplier: multiplier,
		Divisor:    divisor,
	}
}

// `error` interface implementation
func (e ErrDivideByZero) Error() string {
	return fmt.Sprintf("divide by zero error, multiplier = %d and divisor = %d", e.Multiplier, e.Divisor)
}

func main() {
	// divisor := 7
	divisor := 0
	q, err := divide(100, divisor)
	if err == nil {
		fmt.Printf("Q : %d\n", q)
	} else {
		fmt.Println(err)
	}
}

func divide(multiplier, divisor int) (quotient int, err error) {
	if divisor == 0 {
		err = NewErrDivideByZero(multiplier, divisor)
		return
	}
	quotient = multiplier / divisor
	return
}
