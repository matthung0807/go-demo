package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(rw http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name") // get URL query string
		content := fmt.Sprintf("hello, %s", name)
		fmt.Fprint(rw, content) // write out content
	})

	http.ListenAndServe(":8080", nil)
}
