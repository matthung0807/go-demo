package main

import (
	"fmt"

	"github.com/samber/lo"
)

func main() {
	map1 := map[string]int{
		"john": 33,
		"mary": 28,
	}
	map2 := map[string]int{
		"mary": 18,
		"tony": 40,
	}
	resultMap := lo.Assign(map1, map2)
	fmt.Println(resultMap)
}
