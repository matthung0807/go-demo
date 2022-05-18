package main

import (
	"fmt"

	"abc.com/demo/db"
	"abc.com/demo/repo"
)

func main() {
	db := db.OpenDB()
	defer db.Close()

	var er repo.EmployeeRepository
	er = repo.NewEmployeeRepository(db)
	emps, err := er.GetAllEmployees()
	if err != nil {
		panic(err)
	}
	fmt.Println(emps)

	emp, err := er.GetEmployeeByID(db, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(*emp)
}
