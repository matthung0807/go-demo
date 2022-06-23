package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestEmployeeHandler_Post(t *testing.T) {
	// create testing request and response writer
	w := httptest.NewRecorder()
	body := "{\"id\":1,\"name\":\"John\",\"age\":33}"
	r := httptest.NewRequest(http.MethodPost, "/employee", strings.NewReader(body))

	// send testing request and response writer to target handler
	EmployeeHandler(w, r)

	// read response body
	b, err := io.ReadAll(w.Body)
	if err != nil {
		t.Errorf("read response body error, err=%v", err)
	}
	// verify result as expected
	result := string(b)
	if result != body {
		t.Errorf("expected %v, but %v", body, result)
	}
}

func TestEmployeeEndpoint_Post(t *testing.T) {
	// create mock server with target endpoint handler
	server := httptest.NewServer(http.HandlerFunc(EmployeeHandler))
	defer server.Close()

	// create http request with mock server's url
	body := "{\"id\":1,\"name\":\"John\",\"age\":33}"
	req, err := http.NewRequest(http.MethodPost, server.URL, strings.NewReader(body))
	if err != nil {
		t.Errorf("create request error, err=%v", err)
	}

	// send http request to mock server
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("send request error, err=%v", err)
	}

	// read response body
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("read response body error, err=%v", err)
	}
	defer resp.Body.Close()

	// verify result as expected
	result := string(b)
	if result != body {
		t.Errorf("expected %v, but %v", body, result)
	}
}
