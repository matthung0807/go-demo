package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	amqp "github.com/rabbitmq/amqp091-go"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // diable CORS check
	},
}

func main() {

	http.HandleFunc("/notification", notificationHandler)
	http.ListenAndServe(":8080", nil)
}

func notificationHandler(w http.ResponseWriter, r *http.Request) {
	mqConn, err := amqp.Dial("amqp://guest:guest@localhost:5672/") // create rabbitmq connection
	if err != nil {
		panic(err)
	}
	defer mqConn.Close()

	ch, err := mqConn.Channel() // create message channel
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	// consume message from queue 'hello'
	msgs, err := ch.Consume(
		"hello",          // queue name
		"hello-consumer", // consumer name
		true,             // auto-ack
		false,            // exclusive
		false,            // no-local
		false,            // no-wait
		nil,              // args
	)
	if err != nil {
		panic(err)
	}

	wsConn, err := upgrader.Upgrade(w, r, nil) // get a websocket connection
	if err != nil {
		panic(err)
	}
	defer wsConn.Close()
	for {
		messageHandler := func(bytes []byte) error {
			fmt.Printf("push message=\"%s\"\n", bytes)
			err = wsConn.WriteMessage(websocket.TextMessage, bytes) // write a message to client
			if err != nil {
				return err
			}
			return nil
		}
		var forever chan struct{}
		receive(msgs, messageHandler) // pass consumed message to messageHandler
		<-forever
	}
}

type MessageHandler func(bytes []byte) error

func receive(msgs <-chan amqp.Delivery, handler MessageHandler) {
	go func() {
		for d := range msgs {
			handler(d.Body)
		}
	}()
}
