package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Employee struct {
	Name  string
	Age   int
	Photo File
}

type File struct {
	Filename string
	Data     string
}

func main() {
	http.HandleFunc("/employee", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			var emp Employee
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&emp) // decode JSON to Employee
			if err != nil {
				panic(err)
			}

			data, err := base64.StdEncoding.DecodeString(emp.Photo.Data) // decode base64 encoded file data string
			if err != nil {
				panic(err)
			}

			err = os.WriteFile(emp.Photo.Filename, data, 0666) // write file to project root
			if err != nil {
				panic(err)
			}

			fmt.Fprint(w, fmt.Sprintf(
				"name=%v\nage=%v\nphoto=%v",
				emp.Name, emp.Age, emp.Photo.Filename))
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", nil)
}
