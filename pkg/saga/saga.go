package saga

import (
	"context"
	"log"

	"abc.com/demo/internal/event/model"
	"github.com/looplab/fsm"
)

type Saga struct {
	id    string
	name  string
	steps []Step
	SagaState
}

func NewSaga(id, name string, events Events, callbacks fsm.Callbacks) *Saga {
	return &Saga{
		id:        id,
		name:      name,
		SagaState: NewSagaState(events, callbacks),
	}
}

type Action func() (model.Topic, error)
type Compen func() (model.Topic, error)

var Skip = func() (model.Topic, error) { return model.SKIP_TOPIC, nil }

type Step struct {
	action Action
	compen Compen
}

func (s *Saga) AddStep(action Action, compen Compen) *Saga {
	step := Step{action, compen}
	s.steps = append(s.steps, step)
	return s
}

func (s *Saga) Execute(ctx context.Context) error {
	log.Printf("saga=[%s] start execute", s.name)
	for _, step := range s.steps {
		topic, err := step.action()
		if err != nil {
			log.Printf("saga action error, err=[%s]", err)
			return s.Compensate(ctx)
		}
		s.UpdateState(ctx, topic)
		log.Printf("saga=[%s] current state=[%s]", s.name, s.GetCurrentState())
	}
	return nil
}

func (s *Saga) Compensate(ctx context.Context) error {
	log.Printf("saga=[%s] start compensate", s.name)
	for i := len(s.steps) - 1; i >= 0; i-- {
		topic, err := s.steps[i].compen()
		if err != nil {
			log.Printf("saga compensate error, err=[%s]", err)
			return err
		}
		s.UpdateState(ctx, topic)
		log.Printf("saga=[%s] current state=[%s]", s.name, s.GetCurrentState())
	}
	return nil
}

func (s *Saga) GetName() string {
	return s.name
}

func (s *Saga) GetId() string {
	return s.id
}
