package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error catch from panic by recover:", r)
		}
	}()

	var input string
	fmt.Print("Type some number: ")
	fmt.Scanln(&input)

	number, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Error to convert to number, cause:", err.Error())
		return
	}

	fmt.Printf("%d bertipe %T\n", number, number)

	// custom error
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("hello", name)
	} else {
		panic(err.Error())
	}
}

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("Cannot be empty")
	}

	return true, nil
}
