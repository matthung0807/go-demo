package demo

import (
	"testing"

	"abc.com/demo/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type CalculatorMock struct {
	mock.Mock
}

func (calMock *CalculatorMock) Plus(x, y int) int {
	args := calMock.Called(x, y)
	return args.Int(0)
}

func TestAddAge(t *testing.T) {

	testCase := struct {
		x        int
		emp      model.Employee
		expected int
	}{
		1,
		model.Employee{Id: 1, Name: "John", Age: 33},
		34,
	}

	calMock := new(CalculatorMock)       // create mock instance
	calMock.On("Plus", 1, 33).Return(34) // setup mock method arguments

	actual, _ := AddAge(testCase.x, testCase.emp, calMock)

	assert.Equal(t, testCase.expected, actual)
}
