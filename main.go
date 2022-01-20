package main

import (
	"fmt"

	"abc.com/demo/hello"
)

var name string = "main"

func init() {
	fmt.Printf("%s-init-1\n", name)
}

func init() {
	fmt.Printf("%s-init-2\n", name)
}

func init() {
	fmt.Printf("%s-init-3\n", name)
}

func main() {
	hello.Hello()
}
