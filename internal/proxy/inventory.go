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

type InventoryProxyService struct {
	messageService mq.MessageService
}

func NewInventoryProxyService(messageService mq.MessageService) *InventoryProxyService {
	return &InventoryProxyService{
		messageService: messageService,
	}
}

func (s *InventoryProxyService) Create(ctx context.Context, inventory domain.Inventory) (domain.Inventory, error) {
	v := ctx.Value("corId")
	corId := ""
	if v != nil {
		corId = v.(string)
	}
	event := model.NewCreateInventoryEvent(corId, model.CreateInventoryPayload{})
	message, err := json.Marshal(event)
	if err != nil {
		return domain.Inventory{}, err
	}
	err = s.messageService.DirectPublish(ctx, mq.DIRECT_EXCHANGE, mq.GENERAL_QUEUE, message)
	if err != nil {
		return domain.Inventory{}, err
	}

	data, err := s.messageService.ConsumeOne(ctx, mq.CREATE_ORDER_SAGA_CONSUMER, mq.CREATE_ORDER_SAGA_QUEUE)
	if err != nil {
		return domain.Inventory{}, errors.New("create inventory failed")
	}
	log.Printf("consumed reply data=[%s]", string(data))
	var ev model.CreateInventoryReplyEvent
	err = json.Unmarshal(data, &ev)
	if err != nil {
		return domain.Inventory{}, err
	}

	if ev.Payload.Result == model.FAIELD {
		return domain.Inventory{}, errors.New("create inventory failed")
	}

	return ev.Payload.Inventory, nil
}

func (s *InventoryProxyService) Update(ctx context.Context, inventory domain.Inventory) (domain.Inventory, error) {
	v := ctx.Value("corId")
	corId := ""
	if v != nil {
		corId = v.(string)
	}
	event := model.NewUpdateInventoryEvent(corId, model.UpdateInventoryPayload{})
	message, err := json.Marshal(event)
	if err != nil {
		return domain.Inventory{}, err
	}
	err = s.messageService.DirectPublish(ctx, mq.DIRECT_EXCHANGE, mq.GENERAL_QUEUE, message)
	if err != nil {
		return domain.Inventory{}, err
	}

	data, err := s.messageService.ConsumeOne(ctx, mq.CREATE_ORDER_SAGA_CONSUMER, mq.CREATE_ORDER_SAGA_QUEUE)
	if err != nil {
		return domain.Inventory{}, errors.New("update inventory failed")
	}
	log.Printf("consumed reply data=[%s]", string(data))
	var ev model.UpdateInventoryReplyEvent
	err = json.Unmarshal(data, &ev)
	if err != nil {
		return domain.Inventory{}, err
	}

	if ev.Payload.Result == model.FAIELD {
		return domain.Inventory{}, errors.New("update inventory failed")
	}

	return ev.Payload.Inventory, nil
}
