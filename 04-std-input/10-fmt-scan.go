package main

import "fmt"

func main() {
	/* var userName string
	fmt.Println("Enter your name :")
	fmt.Scanln(&userName)
	fmt.Printf("Hi %s, Have a nice day!\n", userName)
	*/

	var firstName, lastName string
	fmt.Println("Enter your name :")
	fmt.Scanln(&firstName, &lastName)
	fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
}
