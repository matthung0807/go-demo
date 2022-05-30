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
			resp := Post("http://localhost:8080/employee")
			fmt.Fprint(w, resp)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.ListenAndServe(":8080", nil)
}

func Post(url string) string {
	resp, err := http.Post(
		url,                 // target url
		"application/json",  // content-type
		createRequestBody()) // request body
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close() // 記得關閉resp.Body

	if resp.StatusCode == http.StatusOK {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		body := string(b)
		return body
	}
	return ""
}

func createRequestBody() *bytes.Buffer {
	emp := Employee{
		Id:   1,
		Name: "john",
		Age:  33,
	}

	data, err := json.Marshal(&emp)
	if err != nil {
		panic(err)
	}

	return bytes.NewBuffer(data)
}
