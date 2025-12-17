package main

import "fmt"

/*
Build an interactive calculator

	Display the menu
	1. Add
	2. Subtract
	3. Multiply
	4. Divide
	5. Exit

	Accept the user choice

	if userchoice <= 4 and >= 1
		Accept two operands
		perform the corresponding operation
		print the result
		DISPLAY THE MENU again

	if userchoice == 5
		print "Thank you!"
		exit the application

	else
		print "Invalid choice"
		DISPLAY THE MENU again
*/

func main() {
	var userChoice, x, y, result int
	for {
		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
		fmt.Println("5. Exit")
		fmt.Println("Enter your choice :")
		fmt.Scanln(&userChoice)
		if userChoice < 1 || userChoice > 5 {
			fmt.Println("Invalid choice...")
			continue
		}
		if userChoice == 5 {
			break
		}
		fmt.Println("Enter the operands")
		fmt.Scanln(&x, &y)
		switch userChoice {
		case 1:
			result = x + y
		case 2:
			result = x - y
		case 3:
			result = x * y
		case 4:
			result = x / y
		}
		fmt.Println("Result :", result)
	}
	fmt.Println("Thank you!")
}
