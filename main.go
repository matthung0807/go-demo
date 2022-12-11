package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("process request")
		time.Sleep(time.Second * 10) // delay 10 seconds to simulate unfinished process
		fmt.Fprint(w, "hello")
	})

	srv := &http.Server{
		Addr: ":8080",
	}

	idleConnsClosed := make(chan struct{}) // block channal to make sure all idle connections closed before server closed
	go func() {
		sigint := make(chan os.Signal, 1)   // channel to receive interrupt signal
		signal.Notify(sigint, os.Interrupt) // receive interrupt signal to `sigint` channel
		<-sigint                            // goroutine blocked here until receving interrup signal
		fmt.Println("\nreceived an interrupt siginal")
		err := srv.Shutdown(context.Background()) // graceful shutdown, wait all active connections become idle then close them
		if err != nil {
			fmt.Printf("graceful shutdown error=%v", err)
		}
		fmt.Println("idle connections closed")
		close(idleConnsClosed)
	}()

	err := srv.ListenAndServe()
	if err != http.ErrServerClosed {
		panic(err)
	}
	<-idleConnsClosed // main goroutine blocked here until 'idleConnsClosed' closed.
	fmt.Println("server closed")
}
