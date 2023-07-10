package domain

import "github.com/google/uuid"

type Order struct {
	Id    string
	State string
}

func NewOrder() Order {
	return Order{
		Id:    uuid.NewString(),
		State: "creating",
	}
}
