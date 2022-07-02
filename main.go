package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type ResponseWriterRecord struct {
	StatusCode int
	Body       []byte
	http.ResponseWriter
}

func NewResponseWriterRecord(w http.ResponseWriter) *ResponseWriterRecord {
	return &ResponseWriterRecord{ResponseWriter: w}
}

func (r *ResponseWriterRecord) Header() http.Header {
	return r.ResponseWriter.Header()
}

func (r *ResponseWriterRecord) Write(b []byte) (int, error) {
	r.Body = b
	return r.ResponseWriter.Write(b)
}

func (r *ResponseWriterRecord) WriteHeader(statusCode int) {
	r.StatusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

func ResponseRecordHandlerFunc(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var rw *ResponseWriterRecord = NewResponseWriterRecord(w)
		h.ServeHTTP(rw, r)

		if rw.StatusCode == 0 {
			rw.StatusCode = http.StatusOK
		}

		header, err := json.Marshal(rw.Header())
		if err != nil {
			panic(err)
		}

		log.Printf("response status code=%d, header=%s, body=%s",
			rw.StatusCode, string(header), string(rw.Body))
	}
}

func main() {
	http.HandleFunc("/hello", HellohandlerFunc())
	http.HandleFunc("/hi", HiHandlerFunc())

	handler := ResponseRecordHandlerFunc(http.DefaultServeMux)
	http.ListenAndServe(":8080", handler)
}

func HellohandlerFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello"))
	}
}

func HiHandlerFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("hi"))
	}
}
