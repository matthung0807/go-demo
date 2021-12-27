package main

import (
	"fmt"

	validator "github.com/go-playground/validator/v10"
)

type Employee struct {
	Id        int    `validate:"required"`                               // 必填
	Name      string `validate:"required,min=3"`                         // 必填，字數長度最小3
	Email     string `validate:"required,email"`                         // 必填，email格式
	Age       int    `validate:"max=65,omitempty"`                       // 數值最大65，忽略空值
	CreatedAt string `validate:"datetime=2006-01-02 15:04:05,omitempty"` // 日期格式yyyy-MM-dd hh:mm:ss，忽略空值
}

func main() {
	emp := Employee{
		Id:        1,
		Name:      "JJ",
		Email:     "JJ.abc.com",
		Age:       70,
		CreatedAt: "2021-12-27 22:12:45",
	}
	err := validator.New().Struct(emp)
	if err != nil {
		fmt.Println(err)
	}
}
