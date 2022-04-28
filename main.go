package main

import (
	"encoding/json"
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
		case http.MethodGet:
			emp := Employee{
				Id:   1,
				Name: "john",
				Age:  33,
			}
			rw.Header().Set("Content-Type", "application/json")
			json.NewEncoder(rw).Encode(emp)
		default:
			http.Error(rw, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.ListenAndServe(":8080", nil)
}
