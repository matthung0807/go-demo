package a

import "fmt"

func A(i int) int {
	fmt.Println(i)
	if i < 1 {
		fmt.Println("i < 1")
	}
	return i
}
