package main

import (
	"net/http"

	c "abc.com/demo/client"
	s "abc.com/demo/server"
)

func main() {
	server()
	client()
	http.ListenAndServe(":8080", nil)
}

func server() {
	s.Route(s.NewWebhooksHandler(s.NewWebhooksService()))
}

func client() {
	c.Route()
}
