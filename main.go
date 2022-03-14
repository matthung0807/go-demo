package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync/atomic"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
	"github.com/aws/aws-sdk-go-v2/service/firehose/types"
)

type Employee struct {
	ID   int64
	Name string
	Age  int
}

func main() {
	http.HandleFunc("/employee", func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var emp Employee
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&emp)
			if err != nil {
				http.Error(rw, "error", http.StatusBadRequest)
			}

			id := Save(&emp)

			data := fmt.Sprintf("employee.id=%d created\n", id)
			go PutRecord(context.TODO(), data)

			fmt.Fprint(rw, id)
		} else {
			http.Error(rw, "not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", nil)
}

var count int64

func Save(emp *Employee) int64 {
	fmt.Printf("save employee=%v\n", emp)
	return atomic.AddInt64(&count, 1)
}

func PutRecord(ctx context.Context, data string) {
	fmt.Printf("put data=%s\n", data)
	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion("ap-northeast-1"),
	)
	if err != nil {
		panic(err)
	}

	client := firehose.NewFromConfig(cfg)

	input := CreateInput(data)
	out, err := client.PutRecord(ctx, input)
	if err != nil {
		fmt.Println("put record error")
	}
	fmt.Printf("recordId=%s\n", *out.RecordId)
}

func CreateInput(data string) *firehose.PutRecordInput {
	deliveryStream := "PUT-s3-demo-bucket-202112151320"
	return &firehose.PutRecordInput{
		DeliveryStreamName: &deliveryStream,
		Record: &types.Record{
			Data: []byte(data),
		},
	}
}
