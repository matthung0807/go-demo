package model

import "abc.com/demo/internal/domain"

// clinet event topic
const (
	CREATE_ORDER_TOPIC Topic = "create_order_topic"
	DELETE_ORDER_TOPIC Topic = "delete_order_topic"
)

// server event topic
const (
	CREATE_ORDER_REPLY_TOPIC Topic = "create_order_reply_topic"
	DELETE_ORDER_REPLY_TOPIC Topic = "delete_order_reply_topic"
)

type CreateOrderEvent struct {
	EventMessage[CreateOrderPayload]
}

func NewCreateOrderEvent(corId string, payload CreateOrderPayload) CreateOrderEvent {
	return CreateOrderEvent{
		EventMessage: NewEventMessage(corId, CREATE_ORDER_TOPIC, payload),
	}
}

type CreateOrderPayload struct {
	UserId string
}

type CreateOrderReplyEvent struct {
	EventMessage[CreateOrderReplyPayload]
}

func NewCreateOrderReplyEvent(corId string, payload CreateOrderReplyPayload) CreateOrderReplyEvent {
	return CreateOrderReplyEvent{
		EventMessage: NewEventMessage(corId, CREATE_ORDER_REPLY_TOPIC, payload),
	}
}

type CreateOrderReplyPayload struct {
	Order  domain.Order
	Result ReplyResult
	UserId string
}

type DeleteOrderEvent struct {
	EventMessage[DeleteOrderPayload]
}

func NewDeleteOrderEvent(corId string, payload DeleteOrderPayload) DeleteOrderEvent {
	return DeleteOrderEvent{
		EventMessage: NewEventMessage(corId, DELETE_ORDER_TOPIC, payload),
	}
}

type DeleteOrderPayload struct {
	Id     string
	UserId string
}

type DeleteOrderReplyEvent struct {
	EventMessage[DeleteOrderReplyPayload]
}

func NewDeleteOrderReplyEvent(corId string, payload DeleteOrderReplyPayload) DeleteOrderReplyEvent {
	return DeleteOrderReplyEvent{
		EventMessage: NewEventMessage(corId, DELETE_ORDER_REPLY_TOPIC, payload),
	}
}

type DeleteOrderReplyPayload struct {
	Result ReplyResult
	UserId string
}
