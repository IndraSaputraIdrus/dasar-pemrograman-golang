package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

var apiUrl = "https://jsonplaceholder.typicode.com/posts"

func getPosts() ([]Post, error) {
	var err error
	var client = &http.Client{}
	var data []Post

	request, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func main() {
	var posts, err = getPosts()
	if err != nil {
		panic(err)
	}

	fmt.Println(posts)
}
