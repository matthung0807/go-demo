package service

import (
	"testing"

	"abc.com/demo/model"
	"abc.com/demo/repo"
)

type TodoRepositoryMock struct {
	insertFn    func(todo *repo.Todo) (*repo.Todo, error)
	getByIDFn   func(id int64) (*repo.Todo, error)
	getByPageFn func(page int, size int) ([]repo.Todo, int, error)
	updateFn    func(todo *repo.Todo) (*repo.Todo, error)
	deleteFn    func(id int64) (*repo.Todo, error)
}

func (trMock *TodoRepositoryMock) Insert(todo *repo.Todo) (*repo.Todo, error) {
	return trMock.insertFn(todo)
}

func (trMock *TodoRepositoryMock) GetByID(id int64) (*repo.Todo, error) {
	return trMock.getByIDFn(id)
}

func (trMock *TodoRepositoryMock) GetByPage(page int, size int) ([]repo.Todo, int, error) {
	return trMock.getByPageFn(page, size)
}

func (trMock *TodoRepositoryMock) Update(todo *repo.Todo) (*repo.Todo, error) {
	return trMock.updateFn(todo)
}

func (trMock *TodoRepositoryMock) Delete(id int64) (*repo.Todo, error) {
	return trMock.deleteFn(id)
}

func TestCreateTodo_Success(t *testing.T) {
	mock := &TodoRepositoryMock{}
	mock.insertFn = func(todo *repo.Todo) (*repo.Todo, error) {
		return &repo.Todo{
			ID: todo.ID,
		}, nil
	}

	testCase := struct{ id, expected int64 }{1, 1}

	ts := NewTodoService(mock)
	result, err := ts.CreateTodo(&model.Todo{ID: testCase.id})
	if err != nil {
		t.Errorf("unexpected error, err=%v", err)
	}
	if result.ID != 1 {
		t.Errorf("expected id=%v, but %v", testCase.expected, result.ID)
	}

}
