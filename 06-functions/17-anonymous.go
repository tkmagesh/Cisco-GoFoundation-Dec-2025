package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hi there!")
	}()

	func(userName string) {
		fmt.Printf("Hello %s!\n", userName)
	}("Magesh")

	func(firstName, lastName string) {
		fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
	}("Magesh", "Kuppan")

	greetMsg := func(firstName, lastName string) string {
		return fmt.Sprintf("Hi %s %s, Have a good day!\n", firstName, lastName)
	}("Suresh", "Kannan")
	fmt.Print(greetMsg)

	q, r := func(multiplier, divisor int) (quotient int, remainder int) { // quotient & remainder will by default be initialized with the 'zero' value of 'int' type
		quotient = multiplier / divisor
		remainder = multiplier % divisor
		return
	}(100, 7)
	fmt.Printf("Dividing %d by %d, quotient = %d and remainder = %d\n", 100, 7, q, r)
}
