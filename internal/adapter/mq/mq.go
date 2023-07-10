package mq

import (
	"context"
)

const (
	DIRECT_EXCHANGE = "direct_exchange"
)

const (
	GENERAL_QUEUE           = "general_queue"
	CREATE_ORDER_SAGA_QUEUE = "create_order_saga_queue"
)

const (
	GENERAL_CONSUMER           = "general_consumer"
	CREATE_ORDER_SAGA_CONSUMER = "create_order_saga_consumer"
)

type MessageHandler func(data []byte) error

type MessageService interface {
	DirectPublish(ctx context.Context, exchange, key string, message []byte) error
	Consume(ctx context.Context, consuemr, queue string, handler MessageHandler) error
	ConsumeOne(ctx context.Context, consuemr, queue string) ([]byte, error)
}
