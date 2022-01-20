package main

import (
	"fmt"
)

// hof is Higher Order Function
func hof(s string, callback func(string)) {
	fmt.Println(s)
	callback(s)
}

// callback is Callback Function
func callback(s string) {
	fmt.Printf("callback: %s\n", s)
}

func main() {
	hof("hi", callback)
}
