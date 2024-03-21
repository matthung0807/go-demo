package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	msgs := genMessages() // 產生多筆訊息

	var wg sync.WaitGroup // 建立WaitGroup物件

	// 每一筆訊息都由一個goroutine處理
	for _, msg := range msgs {
		wg.Add(1) // 每多一個goroutine，就在WaitGroup計數加一
		go func(msg string) {
			defer wg.Done()             // 當一個goroutine結束，就從WaitGroup計數扣除
			time.Sleep(time.Second * 1) // 模擬處理每筆訊息要耗費的時間
			fmt.Println(msg)
		}(msg)
	}
	wg.Wait() // 當WaitGroup的計數為0之前，會阻塞。
}

func genMessages() []string {
	msgs := make([]string, 0)
	for i := 0; i < 10; i++ {
		msgs = append(msgs, fmt.Sprintf("message %d", i))
	}
	return msgs
}
