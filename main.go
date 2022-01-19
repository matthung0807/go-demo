package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Employee struct {
	Id   int
	Name string
	Age  int
}

func main() {
	http.HandleFunc("/employee", func(rw http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			var emp Employee
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&emp)
			if err != nil {
				panic(err)
			}
			fmt.Println(emp)
			json.NewEncoder(rw).Encode(emp)
		default:
			http.Error(rw, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", nil)
}
