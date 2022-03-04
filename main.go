package main

import (
	"fmt"
	"net/http"

	_ "abc.com/demo/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @contact.name   菜鳥工程師肉豬
// @contact.url    https://matthung0807.blogspot.com/
// @title Swagger Demo
// @version 1.0
// @description Swagger API.
// @host localhost:8080
func main() {
	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	http.ListenAndServe(":8080", nil)
}

// @Tags Hello
// @Param name query string false "user name"
// @Success 200 {string} string
// @Router /hello [get]
func HelloHandler(rw http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	content := fmt.Sprintf("hello, %s", name)
	fmt.Fprint(rw, content)
}
