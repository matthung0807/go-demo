package order

import (
	"context"
	"log"

	"abc.com/demo/internal/adapter/repo"
	"abc.com/demo/internal/domain"
	"abc.com/demo/internal/event/model"
	sagaorder "abc.com/demo/internal/saga/order"
	"abc.com/demo/internal/service"
	"abc.com/demo/pkg/saga"
)

type CreateUseCase struct {
	orderService     service.OrderService
	inventoryService service.InventoryService
	sagaRepo         repo.SagaRepo
}

func NewCreateUseCase(
	orderService service.OrderService,
	inventoryService service.InventoryService,
	sagaRepo repo.SagaRepo,
) *CreateUseCase {
	return &CreateUseCase{
		orderService:     orderService,
		inventoryService: inventoryService,
		sagaRepo:         sagaRepo,
	}
}

func (uc *CreateUseCase) Exec(ctx context.Context) error {
	v := ctx.Value("corId")
	corId := ""
	if v != nil {
		corId = v.(string)
	}
	var order domain.Order = domain.NewOrder()
	// var inventory domain.Inventory
	var err error
	createOrderSaga := sagaorder.NewCreateOrderSaga(corId)
	createOrderSaga.AddStep(
		saga.Action(func() (model.Topic, error) {
			err = uc.orderService.Create(ctx, domain.Order{})
			return model.CREATE_ORDER_TOPIC, err
		}),
		saga.Skip,
	).AddStep(
		saga.Action(func() (model.Topic, error) {
			_, err = uc.inventoryService.Update(ctx, domain.Inventory{})
			return model.UPDATE_INVENTORY_TOPIC, err
		}),
		saga.Compen(func() (model.Topic, error) {
			err = uc.orderService.Delete(ctx, order.Id)
			return model.DELETE_ORDER_TOPIC, err
		}),
	)

	err = createOrderSaga.Execute(ctx)
	if err != nil {
		log.Printf("err=[%s]", err)
	}
	sagaDomain := domain.NewSaga(createOrderSaga)
	uc.sagaRepo.Insert(ctx, sagaDomain)

	return nil
}
