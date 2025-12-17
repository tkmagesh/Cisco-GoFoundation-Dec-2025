package main

import "fmt"

func main() {
	sayHi()
	sayHello("Magesh")
	greet("Magesh", "Kuppan")

	greetMsg := getGreetMsg("Suresh", "Kannan")
	fmt.Print(greetMsg)
}

/* no parameters & no return results */
func sayHi() {
	fmt.Println("Hi there!")
}

/* 1 parameter */
func sayHello(userName string) {
	fmt.Printf("Hello %s!\n", userName)
}

/* 2 parameters */
/*
func greet(firstName string, lastName string) {
	fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
}
*/

func greet(firstName, lastName string) {
	fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
}

/* 2 parameters & 1 return result */
func getGreetMsg(firstName, lastName string) string {
	return fmt.Sprintf("Hi %s %s, Have a good day!\n", firstName, lastName)
}
