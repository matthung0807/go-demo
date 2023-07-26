package domain

type Saga struct {
	Id     string // Correlation id
	Name   string // Saga name
	Action string // Saga action
	State  string // Saga state
}

type SagaInterface interface {
	GetId() string
	GetName() string
	GetCurrentState() string
	GetAction() string
}

func NewSaga(si SagaInterface) Saga {
	return Saga{
		Id:     si.GetId(),
		Name:   si.GetName(),
		Action: si.GetAction(),
		State:  si.GetCurrentState(),
	}
}
