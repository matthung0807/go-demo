package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/directconnect"
)

func main() {
	ctx := context.TODO()
	client := NewDirectConnectClient(ctx)

	connectionId := "dxcon-fg5kq63s"
	connectionName := "connection-001"
	params := &directconnect.UpdateConnectionInput{
		ConnectionId:   &connectionId,
		ConnectionName: &connectionName,
	}

	output, err := client.UpdateConnection(ctx, params)
	if err != nil {
		panic(err)
	}
	fmt.Println(*output.ConnectionId)   // dxcon-fg5kq63
	fmt.Println(*output.ConnectionName) // connection-001
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
