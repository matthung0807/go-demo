package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			ctx := context.Background()
			data, err := Get(ctx, "http://localhost:8080/get")
			if err != nil {
				fmt.Fprintf(w, "err=%v", err)
			}
			fmt.Fprintf(w, data)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.ListenAndServe(":8080", nil)
}

func Get(ctx context.Context, url string) (string, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		url,
		nil)
	if err != nil {
		log.Printf("create request error, err=%v", err)
		return "", err
	}

	client := &http.Client{
		Timeout: time.Second * 3,
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("client send request error, err=%v", err)
		return "", err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("read request body error, err=%v", err)
		return "", err
	}
	return string(b), nil
}
