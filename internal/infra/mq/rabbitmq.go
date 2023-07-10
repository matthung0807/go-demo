package mq

import (
	"context"
	"log"

	"abc.com/demo/internal/adapter/mq"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	conn *amqp.Connection
}

func NewRabbitMQ(conn *amqp.Connection) *RabbitMQ {
	return &RabbitMQ{
		conn: conn,
	}
}

func (r *RabbitMQ) DirectPublish(ctx context.Context, exchange, key string, message []byte) error {
	ch, err := r.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		exchange,
		amqp.ExchangeDirect,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	q, err := ch.QueueDeclare(
		key,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	err = ch.QueueBind(
		q.Name,
		key,
		exchange,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	err = ch.PublishWithContext(
		ctx,
		exchange,
		key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *RabbitMQ) Consume(
	ctx context.Context,
	consuemr string,
	queue string,
	handler mq.MessageHandler,
) error {
	ch, err := r.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		queue,    // queue name
		consuemr, // consumer name
		true,     // auto-ack
		false,    // exclusive
		false,    // no-local
		false,    // no-wait
		nil,      // args
	)
	if err != nil {
		return err
	}

	for d := range msgs {
		err = handler(d.Body)
		if err != nil {
			log.Printf("Consume message error, body=[%s], err=[%s]", string(d.Body), err)
		}
	}

	return nil
}

func (r *RabbitMQ) ConsumeOne(ctx context.Context,
	consuemr string,
	queue string,
) ([]byte, error) {
	ch, err := r.conn.Channel()
	if err != nil {
		return nil, err
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		queue,    // queue name
		consuemr, // consumer name
		true,     // auto-ack
		false,    // exclusive
		false,    // no-local
		false,    // no-wait
		nil,      // args
	)
	if err != nil {
		return nil, err
	}

	d := <-msgs
	return d.Body, nil
}
