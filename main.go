package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func main() {
	sch := gocron.NewScheduler(time.UTC) // create a new job scheduler
	count := 0
	_, err := sch.Every(1).Seconds().Do(func() { // create a job to run function every 1 second
		count++
		fmt.Println(count)
	})
	if err != nil {
		panic(err)
	}
	sch.StartAsync()            // start the scheduler's jobs in another goroutine
	time.Sleep(time.Second * 5) // main goroutine sleep for 5 second to let scheduler goroutine have time run
}
