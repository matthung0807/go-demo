package model

import (
	"github.com/google/uuid"
)

type Message interface {
	GetMessageId() string
	GetCorId() string
	GetPayload() Payload
}

type EventMessage[T Payload] struct {
	MessageId string
	CorId     string
	Topic     Topic
	Payload   T
}

func NewEventMessage[T Payload](corId string, topic Topic, payload T) EventMessage[T] {
	return EventMessage[T]{
		MessageId: uuid.NewString(),
		CorId:     corId,
		Topic:     topic,
		Payload:   payload,
	}
}

type Payload interface {
}

func (e *EventMessage[T]) GetMessageId() string {
	return e.MessageId
}
func (e *EventMessage[T]) GetCorId() string {
	return e.CorId
}
func (e *EventMessage[T]) GetTopic() Topic {
	return e.Topic
}
func (e *EventMessage[T]) GetPayload() Payload {
	return e.Payload
}

type ReplyResult string

const (
	SUCCESS ReplyResult = "success"
	FAIELD  ReplyResult = "failed"
)
