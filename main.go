package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var df = "2006-01-02"

type CreatedDate time.Time

func (c *CreatedDate) UnmarshalText(text []byte) error {
	t, err := time.Parse(df, string(text))
	if err != nil {
		return err
	}
	*c = CreatedDate(t)
	return nil
}

func (c CreatedDate) MarshalText() (text []byte, err error) {
	s := time.Time(c).Format(df)
	return []byte(s), nil
}

func (c CreatedDate) String() string {
	return time.Time(c).Format(df)
}

type Employee struct {
	Id   int
	Name string
	Age  int
	CreatedDate
}

func main() {
	http.HandleFunc("/employee", func(rw http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			var emp Employee
			if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
				panic(err)
			}
			fmt.Println(emp.CreatedDate) // 2021-01-19
			json.NewEncoder(rw).Encode(emp)
		default:
			http.Error(rw, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", nil)
}
