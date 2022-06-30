package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Employee struct {
	ID   int64
	Name string
	Age  int
}

func main() {
	http.HandleFunc("/employee", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			var emp Employee
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&emp)
			if err != nil {
				panic(err)
			}
			fmt.Println(emp)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(emp)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", nil)
}
