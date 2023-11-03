package main

import (
	"fmt"

	"go.uber.org/fx"
)

func NewString1() string {
	s := string("foo")
	return s
}

func NewString2() string {
	s := string("bar")
	return s
}

func main() {
	fx.New(
		fx.Provide(
			fx.Annotate(
				NewString1,
				fx.ResultTags(`name:"s1"`),
			),
		),
		fx.Provide(
			fx.Annotate(
				NewString2,
				fx.ResultTags(`name:"s2"`),
			),
		),
		fx.Provide(
			fx.Annotate(
				NewA,
				fx.ParamTags(`name:"s1"`, `name:"s2"`),
			),
		),
		fx.Invoke(func(a *A) {
			fmt.Println("invoke")
		}),
	).Run()
}

type A struct {
	S1 string
	S2 string
}

func NewA(s1, s2 string) *A {
	fmt.Printf("create A with s1=[%s], s2=[%s]\n", s1, s2)
	return &A{
		S1: s1,
		S2: s2,
	}
}
