package client

import "net/http"

func Route() {
	http.HandleFunc("/hello", HelloHandler)
}
