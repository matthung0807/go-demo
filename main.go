package main

import (
	"fmt"

	"go.uber.org/fx"
)

type S1Param struct {
	fx.In

	S string `name:"s1"`
}

type S1Result struct {
	fx.Out

	S string `name:"s1"`
}

type S2Param struct {
	fx.In

	S string `name:"s2"`
}

type S2Result struct {
	fx.Out

	S string `name:"s2"`
}

func NewString1() S1Result {
	s := string("foo")
	return S1Result{
		S: s,
	}
}

func NewString2() S2Result {
	s := string("bar")
	return S2Result{
		S: s,
	}
}

func main() {
	fx.New(
		fx.Provide(NewString1),
		fx.Provide(NewString2),
		fx.Provide(NewA),
		fx.Invoke(func(a *A) {
			fmt.Println("invoke")
		}),
	).Run()
}

type A struct {
	S1 string
	S2 string
}

func NewA(s1 S1Param, s2 S2Param) *A {
	fmt.Printf("create A with s1=[%s], s2=[%s]\n", s1.S, s2.S)
	return &A{
		S1: s1.S,
		S2: s2.S,
	}
}
