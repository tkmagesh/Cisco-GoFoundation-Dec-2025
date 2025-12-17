package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "EXIT" {
			break
		}
		fmt.Println("echo :", txt)
	}
	fmt.Println("Thank you!")
}
