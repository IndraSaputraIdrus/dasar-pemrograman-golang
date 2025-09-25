package main

import (
	"fmt"
	"os"
	"strings"
)

func orderSomeFood(menu string) {
	defer fmt.Println("Terimakasih, silahkan tunggu")

	if strings.ToLower(menu) == "pizza" {
		fmt.Println("Pilihan yang tepat")
		fmt.Println("Pizza ditempat kami paling enak")
		return
	}

	fmt.Println("Pesanan anda:", menu)
}

func main() {
	orderSomeFood("Burger")
	orderSomeFood("Pizza")

	outProgram := true

	if outProgram {
		os.Exit(1)
	}
}
