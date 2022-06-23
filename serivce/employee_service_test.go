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
func (repoMock *EmployeeRepoImplMock) FindEmployeesAgeGreaterThan(age int) []model.Employee {
	args := repoMock.Called(age)
	return args.Get(0).([]model.Employee)
}

func TestGetSrEmployeeNumbers_Age40(t *testing.T) {

	repoMock := new(EmployeeRepoImplMock)
	repoMock.On("FindEmployeesAgeGreaterThan", 40).
		Return([]model.Employee{
			{ID: 99, Name: "Jack", Age: 70},
		})

	expected := 1

	es := EmployeeSerivceImpl{
		EmpRepo: repoMock,
	}

	actial := es.GetSrEmployeeNumbers(40)
	assert.Equal(t, expected, actial)
}
