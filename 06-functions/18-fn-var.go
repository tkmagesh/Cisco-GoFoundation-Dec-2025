package main

import "fmt"

func main() {
	var sayHi func()

	sayHi = func() {
		fmt.Println("Hi there!")
	}
	sayHi()

	sayHi = func() {
		fmt.Println("Hello there!")
	}

	if sayHi != nil {
		sayHi()
	}

	/*  */
	var sayHello func(string)
	sayHello = func(userName string) {
		fmt.Printf("Hello %s!\n", userName)
	}
	sayHello("Magesh")

	var greet func(string, string)
	greet = func(firstName, lastName string) {
		fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
	}
	greet("Magesh", "Kuppan")

	var getGreetMsg func(string, string) string
	getGreetMsg = func(firstName, lastName string) string {
		return fmt.Sprintf("Hi %s %s, Have a good day!\n", firstName, lastName)
	}
	greetMsg := getGreetMsg("Suresh", "Kannan")
	fmt.Print(greetMsg)

	var divide func(int, int) (int, int)
	divide = func(multiplier, divisor int) (quotient int, remainder int) { // quotient & remainder will by default be initialized with the 'zero' value of 'int' type
		quotient = multiplier / divisor
		remainder = multiplier % divisor
		return
	}
	q, r := divide(100, 7)
	fmt.Printf("Dividing %d by %d, quotient = %d and remainder = %d\n", 100, 7, q, r)
}
