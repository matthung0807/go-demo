package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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
			rw.Header().Set("Content-Type", "application/json")
			json.NewEncoder(rw).Encode(emp)
		default:
			http.Error(rw, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			resp, err := http.Post(
				"http://localhost:8080/employee",
				"application/json",
				bytes.NewBuffer(genData()))
			if err != nil {
				panic(err)
			}

			if resp.StatusCode == http.StatusOK {
				b, err := io.ReadAll(resp.Body)
				if err != nil {
					panic(err)
				}
				body := string(b)
				fmt.Printf("body:%s\n", body)
			}
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.ListenAndServe(":8080", nil)
}

func genData() []byte {
	emp := Employee{
		Id:   1,
		Name: "john",
		Age:  33,
	}

	data, err := json.Marshal(&emp)
	if err != nil {
		panic(err)
	}
	return data
}
