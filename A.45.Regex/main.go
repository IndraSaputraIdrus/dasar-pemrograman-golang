package main

import (
	"fmt"
	"regexp"
)

var text = "banana burger soup"

func basicRegex() {
	var regex, err = regexp.Compile("[a-z]+")

	if err != nil {
		fmt.Println("Error in basicRegex:", err.Error())
		return
	}

	var result1 = regex.FindAllString(text, 2)
	fmt.Println(result1)

	var result2 = regex.FindAllString(text, -1)
	fmt.Println(result2)
}

func matchString() {
	var regex, err = regexp.Compile(`[a-z]+`)

	if err != nil {
		fmt.Println("Error in matchString:", err.Error())
		return
	}

	var isMatch = regex.MatchString(text)
	fmt.Println(isMatch)
}

func findString() {
	var regex, err = regexp.Compile(`[a-z]+`)

	if err != nil {
		fmt.Println("Error in findString:", err.Error())
		return
	}

	var str = regex.FindString(text)
	fmt.Println(str)
}

func findStringIndex() {
	var regex, err = regexp.Compile(`[a-z]+`)

	if err != nil {
		fmt.Println("Error in findStringIndex:", err.Error())
		return
	}

	var i = regex.FindStringIndex(text)
	fmt.Println(i)

	fmt.Println(text[i[0]:i[1]])
}

func findAllString() {
	var regex, err = regexp.Compile(`[a-z]+`)

	if err != nil {
		fmt.Println("Error in findAllString:", err.Error())
		return
	}

	var res = regex.FindAllString(text, -1)
	fmt.Println(res)
}

func replaceAllString() {
	var regex, err = regexp.Compile(`[a-z]+`)

	if err != nil {
		fmt.Println("Error in replaceAllString:", err.Error())
		return
	}

	var res = regex.ReplaceAllString(text, "potato")
	fmt.Println(res)
}

func replaceAllStringFunc() {
	var regex, err = regexp.Compile(`[a-z]+`)

	if err != nil {
		fmt.Println("Error in replaceAllStringFunc:", err.Error())
		return
	}

	var callback = func(i string) string {
		if i == "burger" {
			return "potato"
		}
		return i
	}

	var res = regex.ReplaceAllStringFunc(text, callback)
	fmt.Println(res)
}

func split() {
	// split dengan separator adalah karakter "a" dan/atau "b"
	var regex, err = regexp.Compile(`[a-b]+`)

	if err != nil {
		fmt.Println("Error in split:", err.Error())
		return
	}

	var res = regex.Split(text, -1)
	fmt.Println(res)
}

func main() {
	basicRegex()
	matchString()
	findString()
	findStringIndex()
	findAllString()
	replaceAllString()
	replaceAllStringFunc()
	split()
}
