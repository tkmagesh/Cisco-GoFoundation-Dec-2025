package main

import (
	"fmt"

	"github.com/tkmagesh/cisco-gofoundation-dec-2025/12-modularity-app/calc"
	"github.com/tkmagesh/cisco-gofoundation-dec-2025/12-modularity-app/calc/utils"
)

func run() {
	fmt.Println("application initiated")
	fmt.Println("Add Result :", calc.Add(100, 200))
	fmt.Println("Subtract Result :", calc.Subtract(100, 200))

	fmt.Println("Operation Count :", calc.OpCount())

	fmt.Println("Is 21 even ?:", utils.IsEven(21))
}
