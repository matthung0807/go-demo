package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type CreatedTime time.Time

var df = "2006-01-02"

func (c *CreatedTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse(df, s)
	if err != nil {
		return err
	}
	*c = CreatedTime(t)
	return nil
}

func (c CreatedTime) MarshalJSON() ([]byte, error) {
	t := time.Time(c)
	s := t.Format(df)
	return json.Marshal(s)
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
