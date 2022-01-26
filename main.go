package main

import (
	"fmt"
	"net/http"

	_ "abc.com/demo/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Swagger Demo
// @version 1.0
// @description Swagger REST API.
// @host localhost:8080
func main() {
	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	http.ListenAndServe(":8080", nil)
}

// @Success 200 {string} string
// @Router /demo/hello [get]
func HelloHandler(rw http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	content := fmt.Sprintf("hello, %s", name)
	fmt.Fprint(rw, content)
}
