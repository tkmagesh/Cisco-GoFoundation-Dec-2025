package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() // schedule the execution of f1() through the scheduler to be executed in FUTURE
	f2()

	// poor man's synchronization techniques (DO NOT USE)
	// block the execution of main() so that the scheduler will look for other goroutines scheduled and execute them (cooperative multi-tasking)
	time.Sleep(2 * time.Second)
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
