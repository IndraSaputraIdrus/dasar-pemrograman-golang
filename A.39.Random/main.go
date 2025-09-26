package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	randomizer := rand.New(rand.NewSource(10))
	for i := range 3 {
		number := i + 1
		fmt.Printf("random ke-%d: %d\n", number, randomizer.Int())
	}

	randomizer2 := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	for i := range 3 {
		number := i + 1
		fmt.Printf("random2 ke-%d: %d\n", number, randomizer2.Int())
	}

	fmt.Println("random int", randomizer2.Int())
	fmt.Println("random float32", randomizer2.Float32())
	fmt.Println("random uint", randomizer2.Uint32())
	fmt.Println("random number sampai dengan 5:", randomizer2.Intn(5))

	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	randomString := func(length int) string {
		b := make([]rune, length)
		for i := range b {
			b[i] = letters[randomizer2.Intn(len(letters))]
		}
		return string(b)
	}

	fmt.Println("random string 5 karakter:", randomString(5))
}
