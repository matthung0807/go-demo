package main

import (
	"fmt"
)

type Employee struct {
	Id    string
	Name  string
	Email string
	Phone string
	Age   int
}

func (e Employee) String() string {
	return fmt.Sprintf("{id=%s, name=%s, email=%s, phone=%s, age=%d}",
		e.Id, e.Name, e.Email, e.Phone, e.Age)
}

type Option func(e *Employee)

func (option Option) Apply(e *Employee) {
	option(e)
}

func NewEmployee(id string, options ...Option) *Employee {
	e := &Employee{
		Id: id,
	}
	for _, option := range options {
		option(e)
	}
	return e
}

func WithName(name string) Option {
	return func(e *Employee) {
		e.Name = name
	}
}

func WithEmail(email string) Option {
	return func(e *Employee) {
		e.Email = email
	}
}

func WithPhone(phone string) Option {
	return func(e *Employee) {
		e.Phone = phone
	}
}

func WithAge(age int) Option {
	return func(e *Employee) {
		e.Age = age
	}
}

func main() {
	e := NewEmployee(
		"1",
		WithName("john"),
		WithAge(33),
	)
	fmt.Println(e)

	WithEmail("john@abc.com").Apply(e)
	WithPhone("111-222-3333").Apply(e)
	fmt.Println(e)
}
