package main

import (
	"testing"

	"abc.com/demo/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// define mock type
type CalculatorMock struct {
	mock.Mock
}

// use mock to implments CalculatorService's method
func (calMock *CalculatorMock) Plus(x, y int) int {
	args := calMock.Called(x, y)
	return args.Int(0)
}

func (calMock *CalculatorMock) Minus(x, y int) int {
	args := calMock.Called(x, y)
	return args.Int(0)
}

func TestAddAge(t *testing.T) {

	calMock := new(CalculatorMock)       // create mock instance
	calMock.On("Plus", 1, 33).Return(34) // setup mock method arguments and return value

	testCase := struct {
		x        int
		emp      model.Employee
		expected int
	}{
		1,
		model.Employee{Id: 1, Name: "John", Age: 33},
		34,
	}

	actual, _ := AddAge(testCase.x, testCase.emp, calMock) // pass test arguments and mock to replace real one

	assert.Equal(t, testCase.expected, actual)
}

func TestAgeDiff(t *testing.T) {

	calMock := new(CalculatorMock)
	calMock.On("Minus", 33, 23).Return(10)

	testCase := struct {
		emp1     model.Employee
		emp2     model.Employee
		expected int
	}{
		model.Employee{Id: 1, Name: "John", Age: 33},
		model.Employee{Id: 2, Name: "Mary", Age: 23},
		10,
	}

	actual, _ := AgeDiff(testCase.emp1, testCase.emp2, calMock)

	assert.Equal(t, testCase.expected, actual)
}
