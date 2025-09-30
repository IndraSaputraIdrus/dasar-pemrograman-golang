package main

import (
	"fmt"
	"strings"
)

func main() {
	// contains
	isExist := strings.Contains("john wick", "wick")
	fmt.Println(isExist)

	// Has prefix ( di awali dengan )
	isPrefix1 := strings.HasPrefix("john wick", "jo")
	fmt.Println(isPrefix1)

	isPrefix2 := strings.HasPrefix("john wick", "wi")
	fmt.Println(isPrefix2)

	// Has suffix ( di akhiri dengan)
	isSuffix1 := strings.HasSuffix("john wick", "jo")
	fmt.Println(isSuffix1)

	isSuffix2 := strings.HasSuffix("john wick", "ck")
	fmt.Println(isSuffix2)

	// Count spesific char
	howMany := strings.Count("ethan hunt", "t")
	fmt.Println(howMany)

	// Index
	// Jika parameter ke2 itu ada banyak substring, maka di ambil yang pertama
	indexChar := strings.Index("ethan hunt", "hu")
	fmt.Println(indexChar)

	indexChar2 := strings.Index("ethan hunt", "u")
	fmt.Println(indexChar2)

	// Replace
	// Parameter terakhir itu jumlah replace akan berlaku, jika -1 akan replace semua
	var name = "indra saputra idrus"
	newName := strings.Replace(name, "a", "i", 2)
	fmt.Println(newName)

	// Repeat
	var laughInIndo = strings.Repeat("wk", 5)
	fmt.Println(laughInIndo)

	// Split
	var names = strings.Split(name, " ")
	fmt.Println(names)

	var nameChar = strings.Split(names[0], "")
	fmt.Println(nameChar)

	// Join
	var firstName = strings.Join(nameChar, "-")
	fmt.Println(firstName)

	var randomName = "NaMA aLAy"

	// to lower
	var lowerName = strings.ToLower(randomName)

	// to upper
	var upperName = strings.ToUpper(randomName)

	fmt.Printf("%v | %v", lowerName, upperName)
}
