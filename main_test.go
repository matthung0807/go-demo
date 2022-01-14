package main

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetAllEmployees(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("opening a stub database connection error=%s", err)
	}
	defer db.Close()

	// mock return rows
	rows := sqlmock.NewRows([]string{"id", "name", "age", "created_on"}).
		AddRow(1, "john", 33, time.Now()).
		AddRow(2, "mary", 28, time.Now())

	mock.ExpectQuery("^SELECT (.+) FROM employee$").WillReturnRows(rows)

	emps, err := GetAllEmployees(db)
	count := len(emps)
	expected := 2
	if count != expected {
		t.Errorf("Expected %d, but %d", expected, count)
	}
}

func TestGetEmployeeByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("opening a stub database connection error=%s", err)
	}
	defer db.Close()

	createdOn, _ := time.Parse("2006-01-02", "2021-01-14")
	rows := sqlmock.NewRows([]string{"id", "name", "age", "created_on"}).
		AddRow(1, "john", 33, createdOn)

	// mock return rows
	mock.ExpectQuery("^SELECT (.+) FROM employee WHERE id = \\$1 LIMIT 1$").
		WillReturnRows(rows)

	emp, err := GetEmployeeByID(db, 1)

	expected := Employee{
		ID:        1,
		Name:      "john",
		Age:       33,
		CreatedOn: createdOn,
	}
	if *emp != expected {
		t.Errorf("Expected %v, but %v", expected, emp)
	}
}
