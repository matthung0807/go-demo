package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Date time.Time

var df = "2006-01-02"

func (d *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse(df, s)
	if err != nil {
		return err
	}
	*d = Date(t)
	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	t := time.Time(d)
	s := t.Format(df)
	return json.Marshal(s)
}

type Employee struct {
	Id          int
	Name        string
	Age         int
	CreatedDate Date
}

func main() {
	http.HandleFunc("/employee", func(rw http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			var emp Employee
			if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
				panic(err)
			}
			fmt.Println(time.Time(emp.CreatedDate)) // 2021-01-19 00:00:00 +0000 UTC
			json.NewEncoder(rw).Encode(emp)
		default:
			http.Error(rw, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", nil)
}
