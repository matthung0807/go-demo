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
	http.ListenAndServe(":8080", router())
}

func router() http.Handler {
	ts := todoService()
	router := httprouter.New()
	router.POST("/todo", handler.Create(ts))
	router.GET("/todo/:id", handler.GetByID(ts))
	router.GET("/todo/:page", handler.GetByPage(ts))
	router.PUT("/todo", handler.Update(ts))
	router.DELETE("/todo", handler.Delete(ts))
	return router
}

func todoService() handler.TodoService {
	db := db.OpenDB()
	tr := repo.NewTodoRepository(db)
	ts := service.NewTodoService(tr)
	return ts
}
