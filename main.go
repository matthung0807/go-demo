package main

import (
	"fmt"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(NewA),
		fx.Provide(NewB),
		fx.Invoke(func(a *A) {
			fmt.Println("invoke")
		}),
	).Run()
}

type A struct {
	*B
}

func NewA(b *B) *A {
	fmt.Println("create A")
	return &A{
		B: b,
	}
}

type B struct {
}

func NewB() *B {
	fmt.Println("create B")
	return &B{}
}
