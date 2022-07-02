package main

import (
	"log"
	"net/http"
)

func LoggerHandlerFunc(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s start", r.Method, r.URL.Path)
		h.ServeHTTP(w, r)
		log.Printf("%s %s end", r.Method, r.URL.Path)
	}
}

func main() {
	http.HandleFunc("/hello", LoggerHandlerFunc(HellohandlerFunc()))
	http.HandleFunc("/hi", HiHandlerFunc())

	// handler := LoggerHandlerFunc(http.DefaultServeMux)
	http.ListenAndServe(":8080", nil)
}

func HellohandlerFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := "hello"
		log.Printf("data=%s", data)
		w.Write([]byte(data))
	}
}

func HiHandlerFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := "hi"
		log.Printf("data=%s", data)
		w.Write([]byte(data))
	}
}
