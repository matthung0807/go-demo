package main

import (
	"fmt"

	"abc.com/demo/db"
	"abc.com/demo/model"
	"abc.com/demo/repo"
)

type EmployeeRepository interface {
	GetAllEmployees() ([]model.Employee, error)
	GetEmployeeByID(id int64) (*model.Employee, error)
}

func main() {
	db := db.OpenDB()
	defer db.Close()

	var er EmployeeRepository = repo.NewEmployeeRepository(db)
	emps, err := er.GetAllEmployees()
	if err != nil {
		panic(err)
	}
	fmt.Println(emps)

	emp, err := er.GetEmployeeByID(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(*emp)
}
