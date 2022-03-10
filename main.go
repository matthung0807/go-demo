package main

import "fmt"

type Employee struct {
	ID   int64
	Name string
	Age  int
}

func main() {
	// variable
	i := 3
	fmt.Println(&i)

	// pointer indirection
	var n *int = &i
	fmt.Println((&*n))

	// slice indexing operation
	sl := []int{1, 2, 3}
	fmt.Println(&sl[0])

	// addressable array indexing operation
	arr := [3]string{"a", "b", "c"}
	fmt.Println(&arr[1])

	// addressable struct field selector
	emp := Employee{ID: 1, Name: "John", Age: 33}
	fmt.Println(&emp.ID)

	// composite literals
	fmt.Println(&[]int{1, 2, 3})
	fmt.Println(&map[int]string{1: "a", 2: "b"})
	fmt.Println(&Employee{}) // struct literal is unaddressable, but is legal because of syntax sugar
	// equivalent to tmp := Employee{}; &tmp
}
