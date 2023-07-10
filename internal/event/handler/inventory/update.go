package inventory

import (
	"context"
	"encoding/json"

	"abc.com/demo/internal/adapter/mq"
	"abc.com/demo/internal/domain"
	"abc.com/demo/internal/event/model"
	"abc.com/demo/internal/proxy/reply"
	"abc.com/demo/internal/service"
)

type UpdateInventoryEventHandler struct {
	inventoryService service.InventoryService
	replyService     *reply.UpdateInventoryReplyService
}

func NewUpdateInventoryEventHandler(
	inventoryService service.InventoryService,
	replyService *reply.UpdateInventoryReplyService,
) *UpdateInventoryEventHandler {
	return &UpdateInventoryEventHandler{
		inventoryService: inventoryService,
		replyService:     replyService,
	}
}

func (h *UpdateInventoryEventHandler) Exec(ctx context.Context, data []byte) error {
	var err error
	var inventory domain.Inventory
	defer func() {
		result := model.SUCCESS
		if err != nil {
			result = model.FAIELD
		}
		h.replyService.Reply(ctx, mq.CREATE_ORDER_SAGA_QUEUE, model.UpdateInventoryReplyPayload{
			Inventory: inventory,
			Result:    result,
		})
	}()

	ev := model.UpdateInventoryEvent{}
	err = json.Unmarshal(data, &ev)
	if err != nil {
		return err
	}

	inventory, err = h.inventoryService.Update(ctx, domain.Inventory{
		Id: "inventory-123",
	})
	if err != nil {
		return err
	}

	return nil
}
