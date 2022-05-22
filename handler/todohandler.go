package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"abc.com/demo/model"
	"github.com/julienschmidt/httprouter"
)

type TodoService interface {
	CreateTodo(m *model.Todo) (*model.Todo, error)
	GetTodoById(id int64) (*model.Todo, error)
	GetTodoByPage(page int, size int) (*model.Page, error)
	UpdateTodo(m *model.Todo) (*model.Todo, error)
	DeleteTodo(id int64) (*model.Todo, error)
}

func Create(ts TodoService) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "read request body error", http.StatusBadRequest)
			return
		}
		todo := &model.Todo{}
		err = json.Unmarshal(b, todo)
		if err != nil {
			log.Printf("error=%v\n", err)
			http.Error(w, "unmarshal json error", http.StatusBadRequest)
			return
		}
		result, err := ts.CreateTodo(todo)
		if err != nil {
			log.Printf("error=%v\n", err)
			http.Error(w, "create todo error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&result)
	}
}

func GetByID(ts TodoService) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		id, err := strconv.ParseInt(ps.ByName("id"), 10, 64)
		if err != nil {
			log.Printf("error=%v\n", err)
			http.Error(w, "parse id error", http.StatusBadRequest)
			return
		}
		result, err := ts.GetTodoById(id)
		if err != nil {
			log.Printf("error=%v\n", err)
			http.Error(w, "get todo error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&result)
	}
}

func GetByPage(ts TodoService) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		page, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil {
			log.Printf("error=%v\n", err)
			http.Error(w, "parse page error", http.StatusBadRequest)
			return
		}

		result, err := ts.GetTodoByPage(page, 10)
		if err != nil {
			log.Printf("error=%v\n", err)
			http.Error(w, "get todolist by page error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&result)
	}
}

func Update(ts TodoService) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("error=%v\n", err)
			http.Error(w, "read request body error", http.StatusBadRequest)
			return
		}
		todo := &model.Todo{}
		err = json.Unmarshal(b, todo)
		if err != nil {
			log.Printf("error=%v\n", err)
			http.Error(w, "unmarshal json error", http.StatusBadRequest)
			return
		}
		result, err := ts.UpdateTodo(todo)
		if err != nil {
			log.Printf("error=%v\n", err)
			http.Error(w, "update todo error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&result)
	}
}

func Delete(ts TodoService) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("error=%v\n", err)
			http.Error(w, "read request body error", http.StatusBadRequest)
			return
		}
		todo := &model.Todo{}
		err = json.Unmarshal(b, todo)
		if err != nil {
			log.Printf("error=%v\n", err)
			http.Error(w, "unmarshal json error", http.StatusBadRequest)
			return
		}
		result, err := ts.DeleteTodo(todo.ID)
		if err != nil {
			log.Printf("error=%v\n", err)
			http.Error(w, "delete todo error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&result)
	}
}
