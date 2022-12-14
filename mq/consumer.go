package mq

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consuemr interface {
	Consume(handler MessageHandler) error
}

type MessageHandler func(userId, content string) error

type Message struct {
	UserId  string `json:"userId"`
	Content string `json:"content"`
}

type RabbitMQConsumer struct {
	url string
}

func NewRabbitMQConsumer(url string) RabbitMQConsumer {
	return RabbitMQConsumer{
		url: url,
	}
}

func (rc *RabbitMQConsumer) Consume(queueName, consumerName string, handler MessageHandler) error {
	mqConn, err := amqp.Dial(rc.url) // create rabbitmq connection
	if err != nil {
		return err
	}
	defer mqConn.Close()

	ch, err := mqConn.Channel() // create message channel
	if err != nil {
		return err
	}
	defer ch.Close()

	// consume message from queue 'hello'
	msgs, err := ch.Consume(
		queueName,    // queue name
		consumerName, // consumer name
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	if err != nil {
		return err
	}
	var forever chan struct{}
	go func() {
		for d := range msgs {
			var message Message
			err := json.Unmarshal(d.Body, &message)
			if err != nil {
				log.Printf("unmarshal queue message body=[%v] to message error, err=%v", d.Body, err)
			}
			err = handler(message.UserId, message.Content)
			if err != nil {
				log.Printf("handle userId=[%s]'s queue message error, err=%v", message.UserId, err)
			}
		}
	}()
	<-forever
	return nil
}
