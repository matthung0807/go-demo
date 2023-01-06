package mq

import (
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	mqurl   string
	conn    *amqp.Connection
	channel *amqp.Channel
	IsReady bool
}

func NewRabbitMQ(mqurl string) RabbitMQ {
	return RabbitMQ{
		mqurl: mqurl,
	}
}

func (r *RabbitMQ) Connect(retryDelay time.Duration) {
	go func() {
		for {
			fmt.Println("attempting to connect")
			conn, err := amqp.Dial(r.mqurl)
			if err != nil {
				fmt.Println("failed to connect, retrying...")
				time.Sleep(retryDelay)
				continue
			}
			r.conn = conn

			for {
				fmt.Println("attempting to open channel")
				ch, err := r.conn.Channel()
				if err != nil {
					fmt.Println("failed to open channel, retrying...")
					time.Sleep(retryDelay)
					continue
				}
				r.channel = ch
				r.IsReady = true

				<-r.conn.NotifyClose(make(chan *amqp.Error))
				r.IsReady = false
				fmt.Println("connection closed, reconnecting...")
				break
			}
		}
	}()
}

func (r *RabbitMQ) Consume() {
	for {
		if r.IsReady {
			msgs, err := r.channel.Consume(
				"hello",          // queue name
				"hello-consumer", // consumer
				true,             // auto-ack
				false,            // exclusive
				false,            // no-local
				false,            // no-wait
				nil,              // args
			)
			if err != nil {
				fmt.Println("failed to register a consumer, retrying...")
				continue
			}

			for d := range msgs {
				log.Printf("received a message: %s\n", d.Body)
			}
		}
	}
}
