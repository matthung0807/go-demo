package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type CreatedTime time.Time

func (c *CreatedTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*c = CreatedTime(t)
	return nil
}

func (c CreatedTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(c))
}

type Employee struct {
	Id        int
	Name      string
	Age       int
	CreatedAt CreatedTime
}

func main() {
	http.HandleFunc("/employee", func(rw http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			var emp Employee
			if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
				panic(err)
			}
			fmt.Println(time.Time(emp.CreatedAt))
			json.NewEncoder(rw).Encode(emp)
		default:
			http.Error(rw, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", nil)
}
