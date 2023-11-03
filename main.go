package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func main() {
	ctx := context.TODO()
	client := NewSQSClient(ctx)

	qn := "DemoQueue"
	urlInput := &sqs.GetQueueUrlInput{
		QueueName: &qn,
	}

	urlOutput, err := client.GetQueueUrl(ctx, urlInput)
	if err != nil {
		panic(err)
	}

	msg := "hello world from go"
	smInput := &sqs.SendMessageInput{
		DelaySeconds: 0,
		MessageBody:  &msg,
		QueueUrl:     urlOutput.QueueUrl,
	}

	smOutput, err := client.SendMessage(ctx, smInput)
	if err != nil {
		panic(err)
	}

	fmt.Println(*smOutput.MessageId)
}

func NewSQSClient(ctx context.Context) *sqs.Client {
	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion("ap-northeast-1"),
	)
	if err != nil {
		panic(err)
	}

	return sqs.NewFromConfig(cfg) // Create an Amazon EC2 service client
}
