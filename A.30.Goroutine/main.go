package main

import (
	"fmt"
)

func print(till int, message string) {
	for i := range till {
		fmt.Println((i + 1), message)
	}
}

func main() {
	go print(5, "hello")
	print(5, "world")

	input := ""
	fmt.Scanln(&input)
}
