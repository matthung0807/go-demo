package order

import (
	"context"
	"encoding/json"

	"abc.com/demo/internal/adapter/mq"
	"abc.com/demo/internal/domain"
	"abc.com/demo/internal/event/model"
	"abc.com/demo/internal/proxy/reply"
	"abc.com/demo/internal/service"
)

type CreateOrderEventHandler struct {
	orderService service.OrderService
	replyService *reply.CreateOrderReplyService
}

func NewCreateOrderEventHandler(
	orderService service.OrderService,
	replyService *reply.CreateOrderReplyService,
) *CreateOrderEventHandler {
	return &CreateOrderEventHandler{
		orderService: orderService,
		replyService: replyService,
	}
}

func (h *CreateOrderEventHandler) Exec(ctx context.Context, data []byte) error {
	var err error
	var order domain.Order
	defer func() {
		result := model.SUCCESS
		if err != nil {
			result = model.FAIELD
		}
		h.replyService.Reply(ctx, mq.CREATE_ORDER_SAGA_QUEUE, model.CreateOrderReplyPayload{
			Order:  order,
			Result: result,
		})
	}()

	ev := model.CreateOrderEvent{}
	err = json.Unmarshal(data, &ev)
	if err != nil {
		return err
	}

	err = h.orderService.Create(ctx, domain.Order{
		Id: "order-123",
	})
	if err != nil {
		return err
	}

	return nil
}
