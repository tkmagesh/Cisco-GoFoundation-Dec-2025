package main

import (
	"fmt"
	"time"
)

func main() {
	ch := genNos()
	for no := range ch {
		fmt.Println(no)
	}
	fmt.Println("Done")
}

func genNos() <-chan int {
	ch := make(chan int)
	go func() {
		for i := range 20 {
			time.Sleep(500 * time.Millisecond)
			ch <- i * 10
		}
		close(ch)
	}()
	return ch
}
