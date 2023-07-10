package order

import (
	"abc.com/demo/internal/event/model"
	"abc.com/demo/pkg/saga"
	"github.com/looplab/fsm"
)

const CREATE_ORDER_SAGA = "create_order_saga"

const (
	ORDER_CREATED     saga.State = "order_created"
	INVENTORY_UPDATED saga.State = "inventory_updated"
)

type CreateOrderSaga struct {
	saga.Saga
}

func NewCreateOrderSaga(corId string) *CreateOrderSaga {
	events := saga.Events{
		{Topic: model.CREATE_ORDER_TOPIC, Src: []saga.State{saga.STARTED}, Dst: ORDER_CREATED},
		{Topic: model.DELETE_ORDER_TOPIC, Src: []saga.State{ORDER_CREATED}, Dst: saga.CANCELED},
		{Topic: model.UPDATE_INVENTORY_TOPIC, Src: []saga.State{ORDER_CREATED}, Dst: saga.DONE},
	}
	callbacks := fsm.Callbacks{}

	return &CreateOrderSaga{
		Saga: *saga.NewSaga(corId, CREATE_ORDER_SAGA, events, callbacks),
	}
}

func (c *CreateOrderSaga) GetCurrentState() string {
	return string(c.Saga.GetCurrentState())
}
