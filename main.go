package main

import (
	"fmt"

	validator "github.com/go-playground/validator/v10"
)

type Employee struct {
	Id        int    `validate:"required"`
	Name      string `validate:"required,min=3"`
	Email     string `validate:"required,email"`
	Age       int    `validate:"max=65,omitempty"`
	Contact   string `validate:"json"`
	CreatedAt string `validate:"datetime=2006-01-02 15:04:05,omitempty"`
}

func main() {
	emp := Employee{
		Id:        1,
		Name:      "JJ",
		Email:     "JJ.abc.com",
		Age:       70,
		Contact:   "{\"name\": \"mary\", \"phones\": [\"0912345678\", \"0912654321\"]}",
		CreatedAt: "2021-12-27 22:12:45",
	}
	err := validator.New().Struct(emp)
	if err != nil {
		fmt.Println(err)
	}
}
