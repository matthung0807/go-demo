package main

import (
	"embed"
	"fmt"
)

//go:embed hello.txt
//go:embed tmp/*
var f embed.FS

func main() {
	data, _ := f.ReadFile("hello.txt")
	fmt.Println(string(data)) // Hello world

	data, _ = f.ReadFile("tmp/a.txt")
	fmt.Println(string(data)) // Apple

	data, _ = f.ReadFile("tmp/b.txt")
	fmt.Println(string(data)) // Banana
}
