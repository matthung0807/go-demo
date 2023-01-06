package main

import (
	"time"

	"abc.com/demo/mq"
)

func main() {
	r := mq.NewRabbitMQ("amqp://guest:guest@localhost:5672/")
	r.Connect(time.Second * 5)
	r.Consume()
}
