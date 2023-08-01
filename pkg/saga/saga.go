package saga

import (
	"context"
	"log"

	"github.com/looplab/fsm"
)

type Saga struct {
	id     string
	name   string
	action string
	steps  []Step
	SagaState
}

func NewSaga(id, name, action string, events Events, callbacks fsm.Callbacks) *Saga {
	return &Saga{
		id:        id,
		name:      name,
		action:    action,
		SagaState: NewSagaState(events, callbacks),
	}
}

type Action func() error
type Compen func() error

var Skip = CompenStep{
	Name:   "skip",
	Compen: func() error { return nil },
}

type Step struct {
	action ActionStep
	compen CompenStep
}

type ActionStep struct {
	Name   string
	Action Action
}

type CompenStep struct {
	Name   string
	Compen Compen
}

func (s *Saga) AddStep(action ActionStep, compen CompenStep) *Saga {
	step := Step{action, compen}
	s.steps = append(s.steps, step)
	return s
}

func (s *Saga) Execute(ctx context.Context) error {
	log.Printf("saga=[%s] start execute", s.name)
	for i := 0; i < len(s.steps); i++ {

		action := s.steps[i].action
		s.SetAction(action.Name)
		err := action.Action()
		if err != nil {
			log.Printf("saga action error, err=[%s]", err)
			return s.Compensate(ctx, i)
		}
		s.UpdateState(ctx, action.Name)
		log.Printf("saga=[%s] current state=[%s]", s.name, s.GetCurrentState())
	}
	return nil
}

func (s *Saga) Compensate(ctx context.Context, n int) error {
	log.Printf("saga=[%s] start compensate", s.name)
	for i := n; i >= 0; i-- {
		compen := s.steps[i].compen
		s.SetAction(compen.Name)
		err := compen.Compen()
		if err != nil {
			log.Printf("saga compensate error, err=[%s]", err)
			return err
		}
		s.UpdateState(ctx, compen.Name)
		log.Printf("saga=[%s], action=[%s], current state=[%s]", s.name, s.action, s.GetCurrentState())
	}
	return nil
}

func (s *Saga) GetName() string {
	return s.name
}

func (s *Saga) GetId() string {
	return s.id
}

func (s *Saga) GetAction() string {
	return s.action
}

func (s *Saga) SetAction(action string) {
	s.action = action
}
