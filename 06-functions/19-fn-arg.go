package main

import "fmt"

func main() {
	/*
		exec("F1") // f1 is executed
		exec("f2") // f2 is executed
		exec("non-existent-func")
	*/
	exec(f1)
	exec(f2)
	exec(func() {
		fmt.Println("anonymous function invoked")
	})
}

func exec(fn func()) {
	fn()
}

/*
func exec(fnName string) {
	// invoke either f1 or f2 based on the parameter
	switch fnName {
	case "f1":
		f1()
	case "f2":
		f2()
	default:
		fmt.Println("Invalid argument!")
	}
}
*/

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
