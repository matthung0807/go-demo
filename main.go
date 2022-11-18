package main

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// 取得連線
	conn, err := amqp.Dial("amqp://guest:guest@mq01.mqloud.dev:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// 建立通道
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 建立queue
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// 設定回應時間
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := "hello world" // 訊息內容
	err = ch.PublishWithContext(ctx,
		"",     // default exchange
		q.Name, // routing key
		true,   // mandatory
		false,  // immediate
		// 設定訊息屬性
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf("message sent\n")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
