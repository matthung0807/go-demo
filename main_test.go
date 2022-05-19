package main

import (
	"errors"
	"testing"

	"abc.com/demo/model"
)

type EmployeeRepositoryMock struct {
	getAllEmployeesFn func() ([]model.Employee, error)
	getEmployeeByIDFn func(id int64) (*model.Employee, error)
}

func (erMock *EmployeeRepositoryMock) GetAllEmployees() ([]model.Employee, error) {
	return erMock.getAllEmployeesFn()
}

func (erMock *EmployeeRepositoryMock) GetEmployeeByID(id int64) (*model.Employee, error) {
	return erMock.getEmployeeByIDFn(id)
}

func TestGetEmployeeNumber(t *testing.T) {
	testCase := struct{ mock, expected int }{10, 10}

	erMock := &EmployeeRepositoryMock{}
	erMock.getAllEmployeesFn = func() ([]model.Employee, error) {
		return make([]model.Employee, testCase.mock), nil
	}
	result := GetEmployeeNumber(erMock)
	if result != testCase.expected {
		t.Errorf("expect %v, but %v", testCase.expected, result)
	}
}

func TestGetEmployeeNumber_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expect panic but success")
		}
	}()

	erMock := &EmployeeRepositoryMock{}
	erMock.getAllEmployeesFn = func() ([]model.Employee, error) {
		return nil, errors.New("error")
	}
	GetEmployeeNumber(erMock)
}
