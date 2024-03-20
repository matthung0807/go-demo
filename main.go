package main

import (
	"fmt"
)

func main() {
	messages := genMessages()
	messageChan := make(chan string)
	workerNum := len(messages) / 10
	for i := 0; i < workerNum; i++ {
		go worker(messageChan)
	}
	send(messages, messageChan)
}

func send(messages []string, messageChan chan string) {
	for j := 0; j < len(messages); j++ {
		messageChan <- messages[j]
	}
}

func worker(messageChan chan string) {
	for message := range messageChan {
		fmt.Printf("%s\n", message)
	}
}

func genMessages() []string {
	messages := make([]string, 0)
	for i := 0; i < 100; i++ {
		messages = append(messages, fmt.Sprintf("message %d", i))
	}
	return messages
}
