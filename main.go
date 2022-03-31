package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func main() {

	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion("ap-northeast-1"),
	)
	if err != nil {
		panic(err)
	}

	client := sqs.NewFromConfig(cfg)

	qn := "DemoQueue"
	urlInput := &sqs.GetQueueUrlInput{
		QueueName: &qn,
	}

	urlOutput, err := client.GetQueueUrl(ctx, urlInput)
	if err != nil {
		panic(err)
	}

	rmInput := &sqs.ReceiveMessageInput{
		QueueUrl:            urlOutput.QueueUrl,
		MaxNumberOfMessages: 1,
		VisibilityTimeout:   5,
	}

	rmOutput, err := client.ReceiveMessage(ctx, rmInput)
	if err != nil {
		panic(err)
	}

	if rmOutput.Messages != nil {
		fmt.Println(*rmOutput.Messages[0].MessageId)
		fmt.Println(*rmOutput.Messages[0].Body)
	}
}
