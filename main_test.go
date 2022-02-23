package main

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestUpdateEmployee(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("opening a stub database connection error=%s", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	result := sqlmock.NewResult(0, 1)
	mock.ExpectExec("^UPDATE employee.+$").WillReturnResult(result)
	mock.ExpectCommit()
	err = UpdateEmployee(db, &Employee{})
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
}

func TestUpdateEmployeeReturnID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("opening a stub database connection error=%s", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
	mock.ExpectQuery("^UPDATE employee.+$").WillReturnRows(rows)
	mock.ExpectCommit()
	id, err := UpdateEmployeeReturnID(db, &Employee{})

	expected := int64(1)
	if id != expected {
		t.Errorf("Expected %v, but %v", expected, id)
	}
}
