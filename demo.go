package demo

import (
	"errors"

	"abc.com/demo/model"
	"abc.com/demo/serivce"
)

func AddAge(x int, emp model.Employee, calService serivce.CalculatorService) (int, error) {
	if (emp == model.Employee{}) {
		return -1, errors.New("emp is empty")
	}
	return calService.Plus(x, emp.Age), nil
}
