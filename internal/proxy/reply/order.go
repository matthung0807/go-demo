package reply

import (
	"context"
	"encoding/json"

	"abc.com/demo/internal/adapter/mq"
	"abc.com/demo/internal/event/model"
)

type CreateOrderReplyService struct {
	messageService mq.MessageService
}

func NewCreateOrderReplyService(messageService mq.MessageService) *CreateOrderReplyService {
	return &CreateOrderReplyService{
		messageService: messageService,
	}
}

func (s *CreateOrderReplyService) Reply(ctx context.Context, queue string, payload model.CreateOrderReplyPayload) error {
	v := ctx.Value("corId")
	corId := ""
	if v != nil {
		corId = v.(string)
	}
	event := model.NewCreateOrderReplyEvent(corId, payload)
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

type DeleteOrderReplyService struct {
	messageService mq.MessageService
}

func NewDeleteOrderReplyService(messageService mq.MessageService) *DeleteOrderReplyService {
	return &DeleteOrderReplyService{
		messageService: messageService,
	}
}

func (s *DeleteOrderReplyService) Reply(ctx context.Context, queue string, payload model.DeleteOrderReplyPayload) error {
	v := ctx.Value("corId")
	corId := ""
	if v != nil {
		corId = v.(string)
	}
	event := model.NewDeleteOrderReplyEvent(corId, payload)
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
