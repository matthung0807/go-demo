package saga

import (
	"context"

	"abc.com/demo/internal/event/model"
	"github.com/looplab/fsm"
	"github.com/samber/lo"
)

type State string

const (
	STARTED  State = "started"
	SKIPED   State = "skiped"
	DONE     State = "done"
	CANCELED State = "canceled"
)

type SagaState struct {
	fsm fsm.FSM
}

func NewSagaState(events Events, callbacks fsm.Callbacks) SagaState {
	fsmEvents := lo.Map(events, func(e Event, i int) fsm.EventDesc {
		fmsSrc := lo.Map(e.Src, func(e State, i int) string {
			return string(e)
		})

		return fsm.EventDesc{
			Name: string(e.Topic),
			Src:  fmsSrc,
			Dst:  string(e.Dst),
		}
	})

	return SagaState{
		fsm: *fsm.NewFSM(string(STARTED), fsmEvents, callbacks),
	}
}

func (s *SagaState) UpdateState(ctx context.Context, topic model.Topic) {
	s.fsm.Event(ctx, string(topic))
}

func (s *SagaState) GetCurrentState() State {
	return State(s.fsm.Current())
}

type Event struct {
	Topic model.Topic
	Src   []State
	Dst   State
}
type Events []Event
