package main

import "fmt"

func main() {
	var i int8 = 100 // 1 byte

	var f float64 // 8 byte

	// NO implicity type conversion
	// f = i

	// type conversion syntax (use the type name like a function)
	f = float64(i)

	fmt.Println(f)
}
