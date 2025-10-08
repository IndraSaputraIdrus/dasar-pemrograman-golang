package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Student struct {
	ID    string
	Name  string
	Grade int
}

var data = []Student{
	{"E001", "Ethan", 21},
	{"E002", "John", 25},
	{"E003", "Doe", 24},
	{"E004", "Indra", 22},
	{"E005", "Rafa", 11},
	{"E006", "Adit", 12},
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	var result, err = json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(result)
}

func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	var id = r.URL.Query().Get("id")

	var result []byte
	var err error

	if id == "" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	for _, student := range data {
		if student.ID != id {
			continue
		}

		result, err = json.Marshal(student)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			break
		}
	}

	w.Write(result)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello")
	})

	http.HandleFunc("/users", users)
	http.HandleFunc("/user", user)

	fmt.Println("Server running on port 5173")
	http.ListenAndServe(":5173", nil)
}
