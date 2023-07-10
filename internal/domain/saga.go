package domain

type Saga struct {
	Id    string // Correlation id
	Name  string // Saga name
	State string // Saga state
}

type SagaInterface interface {
	GetId() string
	GetName() string
	GetCurrentState() string
}

func NewSaga(si SagaInterface) Saga {
	return Saga{
		Id:    si.GetId(),
		Name:  si.GetName(),
		State: si.GetCurrentState(),
	}
}
