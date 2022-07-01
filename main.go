package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	h := HelloHandler(HellohandlerFunc)
	mux.Handle("/hello", h)
	// mux.HandleFunc("/hello", HellohandlerFunc)

	http.ListenAndServe(":8080", mux)
}

type HelloHandler func(w http.ResponseWriter, r *http.Request)

func (f HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}

func HellohandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
