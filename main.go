package main

import (
	"fmt"
	"sync"
	"time"
)

var rwmu sync.RWMutex

var m = map[int]int{}

func main() {
	for i := 0; i < 10; i++ {
		go put(i, i)
		go func(k int) {
			fmt.Println(get(k))
		}(i)
	}
	time.Sleep(time.Second)
}

func get(k int) int {
	rwmu.RLock()
	defer rwmu.RUnlock()
	v, ok := m[k]
	if ok == false {
		return -1
	}
	return v
}

func put(k int, v int) {
	rwmu.Lock()
	defer rwmu.Unlock()
	m[k] = v
}

func size() int {
	rwmu.RLock()
	defer rwmu.RUnlock()
	return len(m)
}
