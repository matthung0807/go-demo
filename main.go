package main

import (
	"fmt"
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

func send(messages []string, messageChan chan string) {
	for j := 0; j < len(messages); j++ {
		messageChan <- messages[j] // 將訊息逐筆發送到channel
	}
}

func worker(messageChan chan string) {
	for message := range messageChan { // 將訊息從channel中取出。若channel中無訊息會阻塞
		fmt.Printf("%s\n", message)
	}
}

func genMessages() []string {
	msgs := make([]string, 0)
	for i := 0; i < 100; i++ {
		msgs = append(msgs, fmt.Sprintf("message %d", i))
	}
	return msgs
}
