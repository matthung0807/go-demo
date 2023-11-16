package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/directconnect"
)

func main() {
	ctx := context.TODO()
	client := NewDirectConnectClient(ctx)

	output, err := client.ConfirmConnection(ctx, &directconnect.ConfirmConnectionInput{
		ConnectionId: aws.String("dxcon-fg5kq63s"),
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(output.ConnectionState)
}

func NewDirectConnectClient(ctx context.Context) *directconnect.Client {
	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion("ap-northeast-1"),
	)
	if err != nil {
		panic(err)
	}

	return directconnect.NewFromConfig(cfg) // Create an Amazon Direct Connect service client
}
