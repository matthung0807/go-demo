package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	NewServer()
}

func NewServer() http.Handler {
	router := chi.NewRouter() // create chi router

	router.Get("/employee/{id}", HandHello())

	http.ListenAndServe(":8080", router) // set chi router
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
