package main

import (
	"fmt"
)

func printMessage(what chan string) {
	fmt.Println(<-what)
}

func sayHelloTo2(who string, msg chan<- string) {
	data := fmt.Sprintf("hello %v", who)
	msg <- data
}

func main() {
	message := make(chan string)

	sayHelloTo := func(who string) {
		data := fmt.Sprintf("hello %v", who)
		message <- data
	}

	go sayHelloTo("John doe")
	go sayHelloTo("Peter parker")
	go sayHelloTo("Indra saputra idrus")

	message1 := <-message
	fmt.Println(message1)

	message2 := <-message
	fmt.Println(message2)

	message3 := <-message
	fmt.Println(message3)

	fmt.Println("-----")

	msg := make(chan string)

	names := []string{
		"John",
		"Indra",
		"Peter",
	}

	for _, each := range names {
		go sayHelloTo2(each, msg)
	}

	for range 3 {
		printMessage(msg)
	}
}
