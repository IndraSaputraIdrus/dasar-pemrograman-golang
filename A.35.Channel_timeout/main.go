package main

import (
	"fmt"
	"time"
)

func sendData(ch chan<- int) {
	for i := range 10 {
		ch <- i
		time.Sleep(time.Duration(i+1) * time.Second)
	}
}

func retrieveData(ch <-chan int) {
loop:
	for {
		select {
		case data := <-ch:
			fmt.Printf("Recieve data '%d'\n", data)
		case <-time.After(time.Second * 5):
			fmt.Println("Timeout, no activities under 5 seconds")
			break loop
		}
	}
}

func main() {
	message := make(chan int)

	go sendData(message)
	retrieveData(message)
}
