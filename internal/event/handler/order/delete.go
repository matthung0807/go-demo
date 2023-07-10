package order

import (
	"context"
	"encoding/json"

	"abc.com/demo/internal/adapter/mq"
	"abc.com/demo/internal/event/model"
	"abc.com/demo/internal/proxy/reply"
	"abc.com/demo/internal/service"
)

type DeleteOrderEventHandler struct {
	orderService service.OrderService
	replyService *reply.DeleteOrderReplyService
}

func NewDeleteOrderEventHandler(
	orderService service.OrderService,
	replyService *reply.DeleteOrderReplyService,
) *DeleteOrderEventHandler {
	return &DeleteOrderEventHandler{
		orderService: orderService,
		replyService: replyService,
	}
}

func (h *DeleteOrderEventHandler) Exec(ctx context.Context, data []byte) error {
	var err error
	defer func() {
		result := model.SUCCESS
		if err != nil {
			result = model.FAIELD
		}
		h.replyService.Reply(ctx, mq.CREATE_ORDER_SAGA_QUEUE, model.DeleteOrderReplyPayload{
			Result: result,
		})
	}()

	ev := model.DeleteOrderEvent{}
	err = json.Unmarshal(data, &ev)
	if err != nil {
		return err
	}

	err = h.orderService.Delete(ctx, ev.Payload.Id)
	if err != nil {
		return err
	}

	return nil
}
