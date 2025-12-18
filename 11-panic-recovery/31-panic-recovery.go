package main

import (
	"errors"
	"fmt"
	"os"
)

var ErrDivideByZero error = errors.New("divide by zero error")

func main() {
	defer func() {
		if err := recover(); err != nil {
			// app panicked
			e := fmt.Errorf("app shutdown, cause %w", err.(error))
			fmt.Println("app panicked, error :", e)
			fmt.Println("[graceful shutdown] closing resources")
			os.Exit(22)
		}
		fmt.Println("Thank you!")
	}()
	// divisor := 7
	divisor := 0
	q := divide(100, divisor)
	fmt.Printf("Q : %d\n", q)

}

func divide(multiplier, divisor int) int {
	defer fmt.Println("[divide] - deferred")
	/*
		if divisor == 0 {
			panic(ErrDivideByZero)
		}
	*/
	return multiplier / divisor
}
