package main

import (
	"fmt"

	"abc.com/demo/list"
)

func main() {
	l := list.New()    // []
	l.InsertFirst("b") // [b]
	l.InsertLast("c")  // [b,c]

	fmt.Println(l.First().Value) // b
	fmt.Println(l.Get(0).Value)  // b

	fmt.Println(l.Last().Value)              // c
	fmt.Println(l.Get(1).Value)              // c
	fmt.Println(l.Get(l.Length() - 1).Value) // c

	l.InsertFirst("a") // [a,b,c]

	e := l.First()
	fmt.Println(e.Value)               // a
	fmt.Println(e.Next().Value)        // b
	fmt.Println(e.Next().Next().Value) // c

	l.InsertLast("e")         // [a,b,c,e]
	l.Insert(3, "d")          // [a,b,c,d,e]
	l.Insert(l.Length(), "f") // [a,b,c,d,e,f]

	for node := l.First(); node != nil; node = node.Next() {
		fmt.Print(node.Value) // abcdef
	}

	fmt.Println()
	fmt.Println(l.Length()) // 6

	d := l.Delete(2)      // [a,b,d,e,f]
	fmt.Println(d.Value)  // c
	fmt.Println(e.Next()) // nil

	d = l.Delete(0)       // [b,d,e,f]
	fmt.Println(e.Value)  // a
	fmt.Println(e.Next()) // nil

	d = l.Delete(l.Length() - 1) // [b,d,e]
	fmt.Println(e.Value)         // f
	fmt.Println(e.Next())        // nil

	for node := l.First(); node != nil; node = node.Next() {
		fmt.Print(node.Value) // bde
	}

	fmt.Println()
	fmt.Println(l.Length()) // 3

	l.Clear()               // []
	fmt.Println(l.First())  // nil
	fmt.Println(l.Last())   // nil
	fmt.Println(l.Length()) // 0
}
