package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	ctx, cancel := context.WithCancel(context.Background()) // create cancel context

	ch := make(chan int) // make an unbuffered 'ch' channel
	go send(ch)          // send values to 'ch' channel

	exit := make(chan bool) // make an 'exit' channel. channel block begin
	go func(ctx context.Context) {
		for { // infinit loop
			select {
			case <-ctx.Done(): // when context's Done channel is closed by cancel()
				fmt.Println("done")
				exit <- true // send a value to 'exit' channel
				return
			default:
				i := <-ch // receive values from 'ch' channel
				fmt.Println(i)
				if i == 3 {
					cancel() // close context's Done channel
				}
			}
		}
	}(ctx) // pass cancel context as a goroutine func's parameter
	<-exit // channel block end, receive value from channel

	close(exit) // close channel
	fmt.Println("end")
}

func send(ch chan int) {
	for i := 1; i < 5; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
	}
}
