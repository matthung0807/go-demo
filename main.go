package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	msgs := genMessages()

	var wg sync.WaitGroup
	workerNum := len(msgs) / 10
	wg.Add(workerNum)

	msgChen := make(chan string)
	for i := 0; i < workerNum; i++ {
		go worker(&wg, msgChen)
	}

	send(msgs, msgChen)
	wg.Wait()
}

func send(msgs []string, msgChen chan string) {
	for j := 0; j < len(msgs); j++ {
		msgChen <- msgs[j]
	}
	close(msgChen)
}

func worker(wg *sync.WaitGroup, msgChen chan string) {
	defer wg.Done()
	for msg := range msgChen {
		time.Sleep(time.Second * 1)
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
