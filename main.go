package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/employee", EmployeeHandler)
	http.ListenAndServe(":8080", nil)
}

func EmployeeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, string(b))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
