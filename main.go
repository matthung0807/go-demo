package main

import "fmt"

type Num interface {
	int | float64 // type constraints
}

type Nums[T Num] []T // slice of Num

func main() {
	// call generic function with type arguments
	fmt.Println(Add[int](1, 2))         // 3
	fmt.Println(Add[float64](1.0, 2.2)) // 3.2
	// call generic function without type arguments
	fmt.Println(Add(1, 2))     // 3
	fmt.Println(Add(1.0, 2.2)) // 3.2

	var ints Nums[int] = []int{1, 2, 3}
	PrintAll[int](ints)

	var floats Nums[float64] = []float64{1.1, 2.2, 3.3}
	PrintAll(floats)

	m := map[string]Nums[int]{
		"a": {1, 2},
		"b": {3, 4},
	}

	PrintMap(m)

}

// func Add[T int | float64](x, y T) N {
// 	return x + y
// }

func Add[T Num](x, y T) T {
	return x + y
}

func PrintAll[T Num](nums Nums[T]) {
	for _, n := range nums {
		fmt.Println(n)
	}
}

func PrintMap[K comparable, T Num](m map[K]Nums[T]) {
	for k, v := range m {
		fmt.Printf("key=%v, value=%v\n", k, v)
	}
}
