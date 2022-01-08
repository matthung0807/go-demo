package main

import (
	"embed"
	"fmt"
)

//go:embed hello.txt
//go:embed tmp/*
var tmp embed.FS

func main() {
	data, _ := tmp.ReadFile("hello.txt")
	fmt.Println(string(data)) // Hello world

	data, _ = tmp.ReadFile("tmp/a.txt")
	fmt.Println(string(data)) // Apple

	data, _ = tmp.ReadFile("tmp/b.txt")
	fmt.Println(string(data)) // Banana
}
