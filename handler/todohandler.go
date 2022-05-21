package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"abc.com/demo/model"
)

type TodoService interface {
	CreateTodo(m *model.Todo) (*model.Todo, error)
	GetTodoById(id int64) (*model.Todo, error)
	GetTodoByPage(page int, size int) (*model.Page, error)
	UpdateTodo(m *model.Todo) (*model.Todo, error)
	DeleteTodo(id int64) (*model.Todo, error)
}
type TodoHandler struct {
	todoService TodoService
}

func NewTodoHandler(service TodoService) *TodoHandler {
	return &TodoHandler{
		todoService: service,
	}
}

func HandleTodo(th *TodoHandler) func(rw http.ResponseWriter, rq *http.Request) {
	return func(rw http.ResponseWriter, rq *http.Request) {
		switch rq.Method {
		case http.MethodGet:
		case http.MethodPost:
			th.Post(rw, rq)
		case http.MethodPut:
		case http.MethodDelete:
		default:
			http.Error(rw, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func (th *TodoHandler) Post(rw http.ResponseWriter, rq *http.Request) {
	var todo model.Todo
	b, err := ioutil.ReadAll(rq.Body)
	if err != nil {
		http.Error(rw, "read request body error", http.StatusMethodNotAllowed)
	}
	json.Unmarshal(b, &todo)
	r, err := th.todoService.CreateTodo(&todo)
	if err != nil {
		http.Error(rw, "create todo error", http.StatusMethodNotAllowed)
	}
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(&r)
}
