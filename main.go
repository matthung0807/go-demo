package main

import (
	"context"
	"fmt"

	"github.com/looplab/fsm"
)

func main() {
	fsm := fsm.NewFSM(
		"closed",
		fsm.Events{
			{Name: "open", Src: []string{"closed"}, Dst: "open"},
			{Name: "close", Src: []string{"open"}, Dst: "closed"},
		},
		fsm.Callbacks{},
	)

	fmt.Println(fsm.Current()) // closed

	err := fsm.Event(context.Background(), "open")
	if err != nil {
		panic(err)
	}

	fmt.Println(fsm.Current()) // open

	err = fsm.Event(context.Background(), "close")
	if err != nil {
		panic(err)
	}

	fmt.Println(fsm.Current()) // closed
}
