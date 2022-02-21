package main

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreateEmployee(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("opening a stub database connection error=%s", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	rows := sqlmock.NewRows([]string{"id"}).AddRow(1) // mock return rows
	mock.ExpectQuery("^INSERT INTO employee.+$").WillReturnRows(rows)
	mock.ExpectCommit()
	id, err := CreateEmployee(db, &Employee{})

	expected := int64(1)
	if id != expected {
		t.Errorf("Expected %v, but %v", expected, id)
	}
}
