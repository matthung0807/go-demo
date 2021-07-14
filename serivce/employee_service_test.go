package serivce

import (
	"testing"

	"abc.com/demo/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// define mock type
type EmployeeRepoImplMock struct {
	mock.Mock
}

// use mock to implments CalculatorService's method
func (empRepoImplMock *EmployeeRepoImplMock) FindEmployeesAgeGreaterThan(age int) []model.Employee {
	args := empRepoImplMock.Called(age)
	return args.Get(0).([]model.Employee)
}

func TestGetSrEmployeeNumbers_Age40(t *testing.T) {

	empRepoImplMock := new(EmployeeRepoImplMock)
	empRepoImplMock.On("FindEmployeesAgeGreaterThan", 40).
		Return([]model.Employee{
			{Id: 99, Name: "Jack", Age: 70},
		})

	expected := 1

	empService := EmployeeSerivceImpl{
		EmpRepo: empRepoImplMock,
	}

	actial := empService.GetSrEmployeeNumbers(40)
	assert.Equal(t, expected, actial)
}
