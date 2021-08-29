package main

import (
	"fmt"

	"abc.com/demo/repo"
	"abc.com/demo/serivce"
)

func main() {
	var es serivce.EmployeeService = &serivce.EmployeeSerivceImpl{
		EmpRepo: &repo.EmployeeRepoImpl{},
	}
	n := es.GetSrEmployeeNumbers(30)
	fmt.Printf("There are %d senior employees\n", n)
}
