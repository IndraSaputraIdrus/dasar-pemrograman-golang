package main

import (
	"fmt"
	"time"
)

func main() {
	message := make(chan int, 3)

	go func() {
		for {
			i := <-message
			fmt.Println("Recieve data", i)
		}
	}()

	for i := range 5 {
		fmt.Println("Send data", i)
		message <- i
	}

	time.Sleep(1 * time.Second)
}
