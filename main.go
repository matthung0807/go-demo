package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
	"github.com/aws/aws-sdk-go-v2/service/firehose/types"
)

func CreateInput(data string) *firehose.PutRecordInput {
	deliveryStream := "PUT-s3-demo-bucket-202112151320" // stream name
	return &firehose.PutRecordInput{
		DeliveryStreamName: &deliveryStream,
		Record: &types.Record{
			Data: []byte(data),
		},
	}
}

func main() {
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion("ap-northeast-1"),
	)
	if err != nil {
		panic(err)
	}

	client := firehose.NewFromConfig(cfg)

	data := "Hello world"
	input := CreateInput(data)
	out, err := client.PutRecord(ctx, input)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(out.RecordId) // 0xc000021ae0

}
