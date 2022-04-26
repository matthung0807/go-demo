package server

import "net/http"

func Route(wh *WebhooksHandler) {
	http.HandleFunc("/register", wh.Register)
	http.HandleFunc("/greeting", wh.Greeting)
}
