package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered", r)
		}
	}()
	doPanic()
	fmt.Println("hello world")
}

func doPanic() {
	fmt.Println("before panic")
	defer fmt.Println("defer after panic")

	panic("panic!!")
}
