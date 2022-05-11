package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
)

type KinesisPutRecordAPI interface {
	PutRecord(ctx context.Context,
		params *kinesis.PutRecordInput,
		optFns ...func(*kinesis.Options)) (*kinesis.PutRecordOutput, error)
}

func MakePutRecord(
	ctx context.Context,
	api KinesisPutRecordAPI,
	input *kinesis.PutRecordInput) (*kinesis.PutRecordOutput, error) {

	return api.PutRecord(ctx, input)
}

func CreateInput(payload string) *kinesis.PutRecordInput {
	stream := "KinesisDataStreamDemo" // stream name
	partition := "demo-001"

	return &kinesis.PutRecordInput{
		Data:         []byte(payload),
		PartitionKey: &partition,
		StreamName:   &stream,
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

	client := kinesis.NewFromConfig(cfg)

	payload := "Hello world"
	input := CreateInput(payload)

	results, err := MakePutRecord(ctx, client, input)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(*results.SequenceNumber)
}
