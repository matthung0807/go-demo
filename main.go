package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getenv("FOO")) // BAR
	fmt.Println(os.Getenv("HI"))  // YO
}
