package proxy

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"abc.com/demo/internal/adapter/mq"
	"abc.com/demo/internal/domain"
	"abc.com/demo/internal/event/model"
)

type OrderProxyService struct {
	messageService mq.MessageService
}

func NewOrderProxyService(messageService mq.MessageService) *OrderProxyService {
	return &OrderProxyService{
		messageService: messageService,
	}
}

func (s *OrderProxyService) Create(ctx context.Context, order domain.Order) error {
	v := ctx.Value("corId")
	corId := ""
	if v != nil {
		corId = v.(string)
	}
	event := model.NewCreateOrderEvent(corId, model.CreateOrderPayload{})
	message, err := json.Marshal(event)
	if err != nil {
		return nil
	}
	err = s.messageService.DirectPublish(ctx, mq.DIRECT_EXCHANGE, mq.GENERAL_QUEUE, message)
	if err != nil {
		return nil
	}

	data, err := s.messageService.ConsumeOne(ctx, mq.CREATE_ORDER_SAGA_CONSUMER, mq.CREATE_ORDER_SAGA_QUEUE)
	if err != nil {
		return errors.New("create order failed")
	}
	log.Printf("consumed reply data=[%s]", string(data))
	var ev model.CreateOrderReplyEvent
	err = json.Unmarshal(data, &ev)
	if err != nil {
		return nil
	}

	if ev.Payload.Result == model.FAIELD {
		return errors.New("create order failed")
	}

	return nil
}

func (s *OrderProxyService) Delete(ctx context.Context, id string) error {
	v := ctx.Value("corId")
	corId := ""
	if v != nil {
		corId = v.(string)
	}
	event := model.NewDeleteOrderEvent(corId, model.DeleteOrderPayload{})
	message, err := json.Marshal(event)
	if err != nil {
		return err
	}
	err = s.messageService.DirectPublish(ctx, mq.DIRECT_EXCHANGE, mq.GENERAL_QUEUE, message)
	if err != nil {
		return err
	}

	data, err := s.messageService.ConsumeOne(ctx, mq.CREATE_ORDER_SAGA_CONSUMER, mq.CREATE_ORDER_SAGA_QUEUE)
	log.Printf("consumed reply data=[%s]", string(data))
	var ev model.DeleteOrderReplyEvent
	err = json.Unmarshal(data, &ev)
	if err != nil {
		return err
	}

	if ev.Payload.Result == model.FAIELD {
		return errors.New("delete order failed")
	}

	return nil
}
