package router

import (
	"context"
	"encoding/json"
	"fmt"

	"abc.com/demo/internal/adapter/mq"
	"abc.com/demo/internal/event/model"
)

type EventRouter struct {
	messageService mq.MessageService
	handlerMap     map[model.Topic]EventHandler
}

func NewEventRouter(messageService mq.MessageService) *EventRouter {
	return &EventRouter{
		messageService: messageService,
		handlerMap:     make(map[model.Topic]EventHandler),
	}
}

func (r *EventRouter) Route(ctx context.Context) {
	go r.messageService.Consume(ctx, mq.GENERAL_CONSUMER, mq.GENERAL_QUEUE, r.MessageHandler)
}

func (r *EventRouter) MessageHandler(data []byte) error {
	// log.Printf("consumed data=[%s]", string(data))
	var em model.EventMessage[any]
	err := json.Unmarshal(data, &em)
	if err != nil {
		return err
	}

	ctx := context.WithValue(context.Background(), "corId", em.GetCorId())

	topic := em.GetTopic()
	h, err := r.GetHanlder(topic)
	if err != nil {
		return err
	}

	err = h.Exec(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

type EventHandler interface {
	Exec(ctx context.Context, data []byte) error
}

func (r *EventRouter) AddHandler(topic model.Topic, handler EventHandler) {
	r.handlerMap[topic] = handler
}

func (r *EventRouter) GetHanlder(topic model.Topic) (EventHandler, error) {
	h, ok := r.handlerMap[topic]
	if !ok {
		return nil, fmt.Errorf("EventHandler not found for Topic=[%s]", string(topic))
	}
	return h, nil
}
