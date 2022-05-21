package main

import (
	"net/http"

	"abc.com/demo/db"
	"abc.com/demo/handler"
	"abc.com/demo/repo"
	"abc.com/demo/service"
)

func main() {

}

func route() {

	tr := repo.NewTodoRepository(db.OpenDB())
	ts := service.NewTodoService(tr)
	th := handler.NewTodoHandler(ts)
	http.HandleFunc("/todo", handler.HandleTodo(th))
}
