/* package declaration (mandatory) */
package main

/* import dependency packages */
import "fmt"

/* package scoped variables / constants / type declarations */

/* package init function */

/* main function */
func main() {
	// print("Hello, World!\n")
	// println("Hello, World!")

	fmt.Print("Hello, World!\n")
	// fmt.Println("Hello, World!")

	// fmt.Printf
	// fmt.Printf("Sum of %d and %d is %d\n", 100, 200, 100+200)
	fmt.Printf("Sum of %0.2f and %0.2f is %0.2f\n", 100.0, 200.0, 100.0+200.0)
}

/* other functions */
