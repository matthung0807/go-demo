package serivce

import "abc.com/demo/repo"

type EmployeeService interface {
	GetSrEmployeeNumbers(int) int
}

type EmployeeSerivceImpl struct {
	EmpRepo repo.EmployeeRepo
}

func (es *EmployeeSerivceImpl) GetSrEmployeeNumbers(age int) int {
	srEmps := es.EmpRepo.FindEmployeesAgeGreaterThan(age)
	return len(srEmps)
}
