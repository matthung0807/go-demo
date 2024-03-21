package main

import (
	"fmt"
	"time"
)

func main() {
	msgs := genMessages()        // 產生多筆訊息
	msgChen := make(chan string) // 建立訊息通道(channel)
	workerNum := len(msgs) / 10  // 要建立的goroutine(worker)數量 = 訊息數量 / 10
	for i := 0; i < workerNum; i++ {
		go worker(msgChen) // 建立goroutine來處理channel的訊息
	}
	send(msgs, msgChen) // 將訊息發送到channel
}

func send(msgs []string, msgChen chan string) {
	for j := 0; j < len(msgs); j++ {
		msgChen <- msgs[j] // 將訊息逐筆發送到channel
	}
	close(msgChen) // 訊息都送到channel後將其關閉
}

func worker(msgChen chan string) {
	for msg := range msgChen { // 將訊息從channel中取出。若channel中無訊息會阻塞
		time.Sleep(time.Second * 1) // 模擬處理每筆訊息要耗費的時間
		fmt.Printf("%s\n", msg)
	}
}

func genMessages() []string {
	msgs := make([]string, 0)
	for i := 0; i < 100; i++ {
		msgs = append(msgs, fmt.Sprintf("message %d", i))
	}
	return msgs
}
