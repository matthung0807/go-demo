package reply

import (
	"context"
	"encoding/json"

	"abc.com/demo/internal/adapter/mq"
	"abc.com/demo/internal/event/model"
)

type CreateInventoryReplyService struct {
	messageService mq.MessageService
}

func NewCreateInventoryReplyService(messageService mq.MessageService) *CreateInventoryReplyService {
	return &CreateInventoryReplyService{
		messageService: messageService,
	}
}

func (s *CreateInventoryReplyService) Reply(ctx context.Context, queue string, payload model.CreateInventoryReplyPayload) error {
	v := ctx.Value("corId")
	corId := ""
	if v != nil {
		corId = v.(string)
	}
	event := model.NewCreateInventoryReplyEvent(corId, payload)
	message, err := json.Marshal(event)
	if err != nil {
		return err
	}
	err = s.messageService.DirectPublish(ctx, mq.DIRECT_EXCHANGE, queue, message)
	if err != nil {
		return err
	}
	return nil
}

type UpdateInventoryReplyService struct {
	messageService mq.MessageService
}

func NewUpdateInventoryReplyService(messageService mq.MessageService) *UpdateInventoryReplyService {
	return &UpdateInventoryReplyService{
		messageService: messageService,
	}
}

func (s *UpdateInventoryReplyService) Reply(ctx context.Context, queue string, payload model.UpdateInventoryReplyPayload) error {
	v := ctx.Value("corId")
	corId := ""
	if v != nil {
		corId = v.(string)
	}
	event := model.NewUpdateInventoryReplyEvent(corId, payload)
	message, err := json.Marshal(event)
	if err != nil {
		return err
	}
	err = s.messageService.DirectPublish(ctx, mq.DIRECT_EXCHANGE, queue, message)
	if err != nil {
		return err
	}
	return nil
}
