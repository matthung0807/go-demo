package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	http.ListenAndServe(":8080", NewRouter()) // start serve
}

func NewRouter() http.Handler {
	router := chi.NewRouter() // create chi router
	router.Get("/employee/{id}", HandHello())
	return router
}

func HandHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id") // get path param {id}

		log.Printf("path=%s\n", r.URL.Path)
		log.Printf("id=%s", id)

		w.Write([]byte(id))
	}
}
