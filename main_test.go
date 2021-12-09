package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {

	srv := NewServer()                                    // start chi router
	r := httptest.NewRequest("GET", "/employee/123", nil) // 建立一個http.Request請求
	w := httptest.NewRecorder()                           // 建立一個ResponseRecorder其實作http.ResponseWriter，用來記錄response狀態

	srv.ServeHTTP(w, r) // 傳入ReponseWriter及Request

	expectedStatus := http.StatusOK // 預期HTTP回應狀態碼
	expectedContent := "123"        // 預期回應內容

	statusCode := w.Result().StatusCode // 實際HTTP回應狀態碼
	content := w.Body.String()          // 實際HTTP回應內容
	defer w.Result().Body.Close()

	if statusCode != expectedStatus {
		t.Errorf("expected code == %d, but %d", expectedStatus, statusCode)
	}
	if content != expectedContent {
		t.Errorf("expected content == %s, but %s", expectedContent, content)
	}

}
