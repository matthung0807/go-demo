package service

import (
	"abc.com/demo/model"
	"abc.com/demo/repo"
)

type TodoRepository interface {
	Insert(todo *repo.Todo) (*repo.Todo, error)
	GetByID(id int64) (*repo.Todo, error)
	GetByPage(page int, size int) ([]repo.Todo, int, error)
	Update(todo *repo.Todo) (*repo.Todo, error)
	Delete(id int64) (*repo.Todo, error)
}
type TodoServiceImpl struct {
	todoRepo TodoRepository
}

func NewTodoService(todoRepo TodoRepository) *TodoServiceImpl {
	return &TodoServiceImpl{
		todoRepo: todoRepo,
	}
}

func (ts *TodoServiceImpl) CreateTodo(m *model.Todo) (*model.Todo, error) {
	e, err := ts.todoRepo.Insert(modelToEntity(m))
	if err != nil {
		return nil, err
	}
	return entityToModel(e), nil
}

func (ts *TodoServiceImpl) GetTodoById(id int64) (*model.Todo, error) {
	e, err := ts.todoRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return entityToModel(e), nil
}

func (ts *TodoServiceImpl) GetTodoByPage(page int, size int) (*model.Page, error) {
	es, total, err := ts.todoRepo.GetByPage(page, size)
	if err != nil {
		return nil, err
	}
	var todolist []model.Todo
	for _, e := range es {
		todolist = append(todolist, *entityToModel(&e))
	}
	p := &model.Page{
		Todolist: todolist,
		Total:    total,
		Pages:    total/size + 1,
	}

	return p, nil
}

func (ts *TodoServiceImpl) UpdateTodo(m *model.Todo) (*model.Todo, error) {
	e, err := ts.todoRepo.Update(modelToEntity(m))
	if err != nil {
		return nil, err
	}
	return entityToModel(e), nil
}

func (ts *TodoServiceImpl) DeleteTodo(id int64) (*model.Todo, error) {
	e, err := ts.todoRepo.Delete(id)
	if err != nil {
		return nil, err
	}
	return entityToModel(e), nil
}

func modelToEntity(m *model.Todo) *repo.Todo {
	return &repo.Todo{
		Description: m.Description,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}

func entityToModel(e *repo.Todo) *model.Todo {
	return &model.Todo{
		Description: e.Description,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
	}
}
