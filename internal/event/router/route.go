package router

import (
	"context"

	"abc.com/demo/internal/event/handler/inventory"
	"abc.com/demo/internal/event/handler/order"
	"abc.com/demo/internal/event/model"
	"abc.com/demo/internal/infra/mq"
	"abc.com/demo/internal/infra/postgres/repo"
	"abc.com/demo/internal/proxy/reply"
	"abc.com/demo/internal/service/impl"
	"gorm.io/gorm"
)

func InitServerEventRouter(ctx context.Context, db *gorm.DB, rabbitmq *mq.RabbitMQ) {
	serverRouter := NewEventRouter(rabbitmq)

	orderRepo := repo.NewOrderRepo(db)
	orderService := impl.NewOrderService(orderRepo)
	createOrderReplyService := reply.NewCreateOrderReplyService(rabbitmq)
	inventoryRepo := repo.NewInventoryRepo(db)
	inventoryService := impl.NewInventoryService(inventoryRepo)
	updateInventoryReplyService := reply.NewUpdateInventoryReplyService(rabbitmq)
	deleteOrderReplyService := reply.NewDeleteOrderReplyService(rabbitmq)
	serverRouter.AddHandler(model.CREATE_ORDER_TOPIC, order.NewCreateOrderEventHandler(orderService, createOrderReplyService))
	serverRouter.AddHandler(model.DELETE_ORDER_TOPIC, order.NewDeleteOrderEventHandler(orderService, deleteOrderReplyService))
	serverRouter.AddHandler(model.UPDATE_INVENTORY_TOPIC, inventory.NewUpdateInventoryEventHandler(inventoryService, updateInventoryReplyService))
	serverRouter.Route(ctx)
}
