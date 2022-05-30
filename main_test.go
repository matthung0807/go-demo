package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPost(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello"))
	}))
	defer ts.Close()

	result := Post(ts.URL)
	if result != "hello" {
		t.Errorf("unexpected result=%v", result)
	}
}
