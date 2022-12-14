package main

import (
	"net/http"

	"abc.com/demo/handler"
	"abc.com/demo/mq"
	"abc.com/demo/ws"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // diable CORS check
	},
}

func main() {
	notifier := ws.NewWebSocketManager()
	notificationHandler := handler.NewNotificationHandler(&notifier)
	consumer := mq.NewRabbitMQConsumer("amqp://guest:guest@localhost:5672/")
	go consumer.Consume("hello", "hello-consumer", notificationHandler.MessageHandler())

	http.HandleFunc("/notification", notificationHandler.Handle())
	http.ListenAndServe(":8080", nil)
}
