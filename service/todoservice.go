package service

import (
	"database/sql"
	"time"

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
	e, err := ts.todoRepo.Insert(toEntity(m))
	if err != nil {
		return nil, err
	}
	return toModel(e), nil
}

func (ts *TodoServiceImpl) GetTodoById(id int64) (*model.Todo, error) {
	e, err := ts.todoRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return toModel(e), nil
}

func (ts *TodoServiceImpl) GetTodoByPage(page int, size int) (*model.Page, error) {
	es, total, err := ts.todoRepo.GetByPage(page, size)
	if err != nil {
		return nil, err
	}
	var todolist []model.Todo
	for _, e := range es {
		todolist = append(todolist, *toModel(&e))
	}
	p := &model.Page{
		Todolist: todolist,
		Total:    total,
		Pages:    total/size + 1,
	}

	return p, nil
}

func (ts *TodoServiceImpl) UpdateTodo(m *model.Todo) (*model.Todo, error) {
	e, err := ts.todoRepo.Update(toEntity(m))
	if err != nil {
		return nil, err
	}
	return toModel(e), nil
}

func (ts *TodoServiceImpl) DeleteTodo(id int64) (*model.Todo, error) {
	e, err := ts.todoRepo.Delete(id)
	if err != nil {
		return nil, err
	}
	return toModel(e), nil
}

func toEntity(m *model.Todo) *repo.Todo {
	return &repo.Todo{
		ID:          m.ID,
		Description: m.Description,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   toNullTime(m.UpdatedAt),
		Deleted:     m.Deleted,
	}
}

func toModel(e *repo.Todo) *model.Todo {
	return &model.Todo{
		ID:          e.ID,
		Description: e.Description,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   toTime(e.UpdatedAt),
		Deleted:     e.Deleted,
	}
}

func toNullTime(t *time.Time) sql.NullTime {
	if t == nil {
		return sql.NullTime{}
	}
	return sql.NullTime{
		Time:  *t,
		Valid: true,
	}
}

func toTime(st sql.NullTime) *time.Time {
	if st.Valid == true {
		return &st.Time
	}
	return nil
}
