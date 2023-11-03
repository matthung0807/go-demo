package main

import (
	"fmt"

	"go.uber.org/fx"
)

type Scheduler interface {
	Run()
}

type CleanupScheduler struct {
}

func NewCleanupScheduler() *CleanupScheduler {
	return &CleanupScheduler{}
}

func (sch CleanupScheduler) Run() {
	fmt.Println("start cleanup...")
}

type SyncScheduler struct {
}

func NewSyncScheduler() *SyncScheduler {
	return &SyncScheduler{}
}

func (sch SyncScheduler) Run() {
	fmt.Println("start sync...")
}

func main() {
	fx.New(
		fx.Provide(
			fx.Annotate(
				NewCleanupScheduler,
				fx.As(new(Scheduler)),
				fx.ResultTags(`group:"scheduler"`),
			),
		),
		fx.Provide(
			fx.Annotate(
				NewSyncScheduler,
				fx.As(new(Scheduler)),
				fx.ResultTags(`group:"scheduler"`),
			),
		),
		fx.Provide(
			fx.Annotate(
				NewSchedulerService,
				fx.ParamTags(`group:"scheduler"`),
			),
		),
		fx.Invoke(func(a *SchedulerService) {
			a.Do()
		}),
	).Run()
}

type SchedulerService struct {
	schedulers []Scheduler
}

func NewSchedulerService(schedulers []Scheduler) *SchedulerService {
	return &SchedulerService{
		schedulers: schedulers,
	}
}

func (svc SchedulerService) Do() {
	for _, sch := range svc.schedulers {
		go sch.Run()
	}
}
