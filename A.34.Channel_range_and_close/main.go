package main

import (
	"fmt"
	"time"
)

func sendMessage(ch chan<- string) {
	for i := range 20 {
		ch <- fmt.Sprintf("data %d", i)
		time.Sleep(500 * time.Millisecond)
	}
	close(ch)
}

func printMessage(ch <-chan string) {
	for message := range ch {
		fmt.Println(message)
	}
}

func main() {
	ch := make(chan string)
	go sendMessage(ch)

	printMessage(ch)
}
