package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGet_Success(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello"))
	}))
	defer ts.Close()

	ctx := context.Background()
	result, err := Get(ctx, ts.URL)
	if err != nil {
		t.Errorf("unexpected error, err=%v", err)
	}
	if result != "hello" {
		t.Errorf("unexpected result=%v", result)
	}
}

func TestGet_ContextDeadlineExceeded(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		time.Sleep(time.Second * 2)
		w.Write([]byte("hello"))
	}))
	defer ts.Close()

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	_, err := Get(ctx, ts.URL)
	if err == nil {
		t.Error("unexpected success")
	}
}

func TestGet_ContextCanceled(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		time.Sleep(time.Second * 2)
		w.Write([]byte("hello"))
	}))
	defer ts.Close()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	cancel()

	_, err := Get(ctx, ts.URL)
	if err == nil {
		t.Error("unexpected success")
	}
}

func TestGet_ClientTimeout(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		time.Sleep(time.Second * 5)
		w.Write([]byte("hello"))
	}))
	defer ts.Close()

	_, err := Get(context.Background(), ts.URL)
	if err == nil {
		t.Error("unexpected success")
	}
}
