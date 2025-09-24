// any atau interface{} merupakan tipe data,
// sehingga cara penggunaannya sama seperti tipe data pada umumnya,
// perbedaannya pada variabel bertipe ini nilainya bisa diisi dengan apapun

package main

import (
	"fmt"
	"strings"
)

type person struct {
	Name string
	Age  int
}

func main() {
	// var secret interface{}
	var secret any

	secret = "indra zzz"
	fmt.Printf("%v, type: %T\n", secret, secret)

	secret = []string{"apple", "samsung", "xiaomi"}
	fmt.Printf("%v, type: %T\n", secret, secret)

	secret = 3.14
	fmt.Printf("%v, type: %T\n", secret, secret)

	var data map[string]any

	data = map[string]any{
		"name":      "ethan hunt",
		"grade":     2,
		"breakfast": [3]string{"apple", "manggo", "banana"},
	}

	fmt.Println(data)

	// casting
	var favorite any

	favorite = 2
	var favNumber = favorite.(int) * 20
	fmt.Println(favNumber)

	favorite = []string{"Mlbb", "Ff", "Hok"}
	var favGame = strings.Join(favorite.([]string), ", ")
	fmt.Println(favGame)

	secret = &person{
		Name: "john doe",
		Age:  27,
	}

	var name = secret.(*person).Name
	fmt.Println(name)

	var person = []map[string]any{
		{"Name": "Wick", "Age": 23},
		{"Name": "Ethan", "Age": 23},
		{"Name": "Bourne", "Age": 22},
	}

	for _, each := range person {
		fmt.Printf("%s age is %d\n", each["Name"], each["Age"])
	}

	var fruits = []any{
		map[string]any{"name": "Strawberry", "total": 10},
		[]string{"maggo", "papaya"},
		"melon",
	}

	for _, each := range fruits {
		fmt.Println(each)
	}
}
