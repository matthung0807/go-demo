package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Employee struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"createdAt"`
}

var s string = `
{
  "id": 1,
  "name": "john",
  "age": 33,
  "createdAt": "2022-01-19T12:34:56Z"
}`

func main() {
	var emp Employee
	json.Unmarshal([]byte(s), &emp) // json to struct
	fmt.Println(emp)

	b, err := json.Marshal(emp) // struct to json
	if err != nil {
		panic(err)
	}
	s := string(b)
	fmt.Println(s)
}
