package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	time.Sleep(1 * time.Second)

	duration := time.Since(start)

	fmt.Println("Time elaplsed in seconds", duration.Seconds())
	fmt.Println("Time elaplsed in minutes", duration.Minutes())
	fmt.Println("Time elaplsed in hours", duration.Hours())

	t1 := time.Now()
	time.Sleep(1 * time.Second)
	t2 := time.Now()

	duration2 := t1.Sub(t2)
	fmt.Println("Time elaplsed in seconds", duration2.Seconds())
	fmt.Println("Time elaplsed in minutes", duration2.Minutes())
	fmt.Println("Time elaplsed in hours", duration2.Hours())

	var n time.Duration = 1
	time.Sleep(n * time.Second)

	// Or

	y := 1
	time.Sleep(time.Duration(y) * time.Second)

	fmt.Println("Complete")

}
