package main

import (
	"net/http"

	"abc.com/demo/db"
	"abc.com/demo/handler"
	"abc.com/demo/repo"
	"abc.com/demo/service"
	"github.com/julienschmidt/httprouter"
)

func main() {
	http.ListenAndServe(":8080", route())
}

func route() http.Handler {
	tr := repo.NewTodoRepository(db.OpenDB())
	ts := service.NewTodoService(tr)
	router := httprouter.New()
	router.POST("/todo", handler.Create(ts))
	router.GET("/todo/:id", handler.GetByID(ts))
	return router
}
