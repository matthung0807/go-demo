package main

import (
	"fmt"

	"abc.com/demo/a"
	"abc.com/demo/b"
)

func main() {
	r1 := a.A(1)
	fmt.Println(r1)

	r2 := b.B(1)
	fmt.Println(r2)
}
