package main

import (
	"fmt"
	"encoding/json"
)

type User struct {
	FullName string `json:"name"`
	Age  int    `json:"age"`
}

func decode() {
	var jsonString = `{"name": "Test", "age": 99}`
	var jsonData = []byte(jsonString)

	var data User

	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Name: %s\n", data.FullName)
	fmt.Printf("Age: %d\n", data.Age)
}

func encode() {
	var object = []User{{"Indra s", 22}, {"M Rafa", 11}}
	var jsonData, err = json.Marshal(object)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var jsonString = string(jsonData)
	fmt.Println(jsonString)
}

func main() {
	decode()
	encode()
}
