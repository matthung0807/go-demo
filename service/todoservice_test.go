package service

import (
	"errors"
	"testing"
	"time"

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

func (mock *TodoRepositoryMock) Insert(todo *repo.Todo) (*repo.Todo, error) {
	return mock.insertFn(todo)
}

func (mock *TodoRepositoryMock) GetByID(id int64) (*repo.Todo, error) {
	return mock.getByIDFn(id)
}

func (mock *TodoRepositoryMock) GetByPage(page int, size int) ([]repo.Todo, int, error) {
	return mock.getByPageFn(page, size)
}

func (mock *TodoRepositoryMock) Update(todo *repo.Todo) (*repo.Todo, error) {
	return mock.updateFn(todo)
}

func (mock *TodoRepositoryMock) Delete(id int64) (*repo.Todo, error) {
	return mock.deleteFn(id)
}

func TestCreateTodo(t *testing.T) {
	testCase := struct{ id, expected int64 }{1, 1}
	mock := &TodoRepositoryMock{
		insertFn: func(todo *repo.Todo) (*repo.Todo, error) {
			return &repo.Todo{
				ID: todo.ID,
			}, nil
		},
	}

	ts := NewTodoService(mock)
	result, err := ts.CreateTodo(&model.Todo{ID: testCase.id})
	if err != nil {
		t.Errorf("unexpected error, err=%v", err)
	}
	if result.ID != 1 {
		t.Errorf("expect id=%v, but %v",
			testCase.expected, result.ID)
	}

}

func TestCreateTodo_Fail(t *testing.T) {
	testCase := struct{ id, expected int64 }{1, 1}
	mock := &TodoRepositoryMock{
		insertFn: func(todo *repo.Todo) (*repo.Todo, error) {
			return nil, errors.New("error")
		},
	}

	ts := NewTodoService(mock)
	_, err := ts.CreateTodo(&model.Todo{ID: testCase.id})
	if err == nil {
		t.Errorf("unexpected success")
	}

}

func TestGetByID(t *testing.T) {
	testCase := struct {
		id       int64
		expected string
	}{
		1, "test",
	}

	mock := &TodoRepositoryMock{
		getByIDFn: func(id int64) (*repo.Todo, error) {
			return &repo.Todo{Description: "test"}, nil
		},
	}
	ts := NewTodoService(mock)
	result, err := ts.GetTodoByID(testCase.id)
	if err != nil {
		t.Errorf("unexpected error, err=%v", err)
	}
	if result.Description != testCase.expected {
		t.Errorf("expect description=%v, but %v",
			testCase.expected, result.Description)
	}

}

func TestGetByID_Fail(t *testing.T) {
	testCase := struct {
		id       int64
		expected string
	}{
		1, "test",
	}
	mock := &TodoRepositoryMock{
		getByIDFn: func(id int64) (*repo.Todo, error) {
			return nil, errors.New("error")
		},
	}

	ts := NewTodoService(mock)
	_, err := ts.GetTodoByID(testCase.id)
	if err == nil {
		t.Errorf("unexpected success")
	}

}

func TestUpdateTodo(t *testing.T) {
	now := time.Now()
	testCase := struct {
		id        int64
		updatedAt time.Time
		expected  time.Time
	}{1, now, now}
	mock := &TodoRepositoryMock{
		updateFn: func(todo *repo.Todo) (*repo.Todo, error) {
			return &repo.Todo{
				ID:        todo.ID,
				UpdatedAt: todo.UpdatedAt,
			}, nil
		},
	}

	ts := NewTodoService(mock)
	result, err := ts.UpdateTodo(&model.Todo{ID: testCase.id})
	if err != nil {
		t.Errorf("unexpected error, err=%v", err)
	}
	if result.ID != 1 {
		t.Errorf("expect updatedAt=%v, but %v",
			testCase.expected, result.UpdatedAt)
	}

}
