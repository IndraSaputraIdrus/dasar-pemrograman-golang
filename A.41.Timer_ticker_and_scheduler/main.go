package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	// fmt.Println("Start")
	// time.Sleep(2 * time.Second)
	// fmt.Println("After 2 seconds")
	//
	// timer2 := time.NewTimer(4 * time.Second)
	// fmt.Println("Start timer")
	// <-timer2.C
	// fmt.Println("Finish timer")
	//
	// ch := make(chan bool)
	// time.AfterFunc(4*time.Second, func() {
	// 	fmt.Println("expired")
	// 	ch <- true
	// })
	//
	// fmt.Println("Start AfterFunc")
	// <-ch
	// fmt.Println("Finish AfterFunc")
	//
	// <-time.After(4 * time.Second)
	// fmt.Println("expired")
	//
	// done := make(chan bool)
	// ticker := time.NewTicker(time.Second)
	//
	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	done <- true
	// }()
	//
	// for {
	// 	select {
	// 	case <-done:
	// 		ticker.Stop()
	// 	case t := <-ticker.C:
	// 		fmt.Println("Hello world", t)
	// 	}
	// }
	//
	timer := func(timeout int, ch chan<- bool) {
		time.AfterFunc(time.Duration(timeout) * time.Second, func () {
			ch <- true
		})
	}

	watcher := func(timeout int, ch <-chan bool) {
		<-ch
		fmt.Printf("\ntimeout! no answer more than %d second", timeout)
		os.Exit(0)
	}

	timeout := 5
	ch2 := make(chan bool)

	go timer(timeout, ch2)
	go watcher(timeout, ch2)

	input := ""
	fmt.Println("what is 725/25 ? ")
	fmt.Scanln(&input)

	if i, _ := strconv.Atoi(input); i == 29 {
		fmt.Println("The answer is right!")
	} else {
		fmt.Println("The answer is wrong!")
	}
}
