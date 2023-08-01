package order

import (
	"abc.com/demo/internal/domain"
	"abc.com/demo/pkg/saga"
	"github.com/looplab/fsm"
)

const CREATE_ORDER_SAGA = "create_order_saga"

const (
	ORDER_CREATED     saga.State = "order_created"
	INVENTORY_UPDATED saga.State = "inventory_updated"
)

const (
	START                   = "start"
	CREATE_ORDER            = "create_order"
	CREATE_ORDER_COMPENSATE = "create_order_compensate"
	UPDATE_INVENTORY        = "update_inventory"
)

type CreateOrderSaga struct {
	saga.Saga
}

func NewCreateOrderSaga(corId string) *CreateOrderSaga {
	events := saga.Events{
		{Name: CREATE_ORDER,
			Src: []saga.State{saga.STARTED},
			Dst: ORDER_CREATED,
		},
		{Name: CREATE_ORDER_COMPENSATE,
			Src: []saga.State{ORDER_CREATED},
			Dst: saga.CANCELED,
		},
		{Name: UPDATE_INVENTORY,
			Src: []saga.State{ORDER_CREATED},
			Dst: saga.DONE,
		},
	}

	callbacks := fsm.Callbacks{}

	return &CreateOrderSaga{
		Saga: *saga.NewSaga(corId, CREATE_ORDER_SAGA, START, events, callbacks),
	}
}

func (c *CreateOrderSaga) GetCurrentState() string {
	return string(c.Saga.GetCurrentState())
}

func NewCreateOrderSagaFromDomain(domainSaga domain.Saga) *CreateOrderSaga {
	createOrderSaga := NewCreateOrderSaga(domainSaga.Id)
	createOrderSaga.SetAction(domainSaga.Action)
	createOrderSaga.SetState(domainSaga.State)
	return createOrderSaga
}
