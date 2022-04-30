package main

import (
	"fmt"
	"sync"
	"time"
)

var count int

var mu sync.Mutex

func main() {
	go inc("a")
	go inc("b")

	time.Sleep(time.Second)
}

func inc(name string) {
	mu.Lock()
	count++
	fmt.Printf("%s:%d\n", name, count)
	mu.Unlock()
}
