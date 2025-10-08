package main

import (
	"fmt"
	"net/url"
)

func main() {
	var myUrl = "http://foo.com/bar?name=test&age=99"
	var u, err = url.Parse(myUrl)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("url: %s\n", myUrl)

	fmt.Printf("protocol: %s\n", u.Scheme)
	fmt.Printf("host: %s\n", u.Host)
	fmt.Printf("path: %s\n", u.Path)

	var name = u.Query()["name"][0]
	var age = u.Query()["age"][0]

	fmt.Printf("name: %s, age: %s", name, age)
}
