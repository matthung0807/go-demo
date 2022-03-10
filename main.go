package main

import "fmt"

type Num int

// func (n *Num) Count() {
// 	*n = *n + 1
// }

func (n Num) Double() int {
	return int(n) * 2
}

type Employee struct {
	ID   int64
	Name string
	Age  int
}

func (emp *Employee) String() string {
	return fmt.Sprintf("ID:%d:%s(%d)", emp.ID, emp.Name, emp.Age)
}

func main() {
	// variable
	i := 3
	fmt.Println(&i)

	// pointer indirection
	var n Num = 10
	fmt.Println((&n).Double())

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
	fmt.Println(&Employee{ID: 2, Name: "Mary", Age: 28})
}
