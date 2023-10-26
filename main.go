package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	b, err := json.Marshal(map[string]string{
		"email":    "john@abc.com",
		"password": "12345",
	})
	if err != nil {
		panic(err)
	}

	body := bytes.NewReader(b)
	url := "http://abc.com/login"
	contentType := "application/json"

	resp, err := http.Post(url, contentType, body) // POST request
	if err != nil {
		panic(err)
	}

	b, err = ioutil.ReadAll(resp.Body) // read response body content
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
