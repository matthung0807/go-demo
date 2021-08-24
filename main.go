package main

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

func AgeDiff(emp1, emp2 model.Employee, calService serivce.CalculatorService) (int, error) {
	if (emp1 == model.Employee{}) || (emp2 == model.Employee{}) {
		return 0, errors.New("one emp is empty")
	}
	return calService.Minus(emp1.Age, emp2.Age), nil
}
