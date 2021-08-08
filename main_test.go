package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	server := setupServer()

	req, _ := http.NewRequest("GET", "/hello", nil) // 建立一個請求
	w := httptest.NewRecorder()                     // 建立一個ResponseRecorder其實作http.ResponseWriter，用來記錄response狀態
	server.ServeHTTP(w, req)                        // gin.Engine.ServerHttp實作http.Handler介面，用來處理HTTP請求及回應。

	expectedStatus := http.StatusOK
	expectedContent := "hello world"

	assert.Equal(t, expectedStatus, w.Code)
	assert.Contains(t, w.Body.String(), expectedContent)
}
