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
	params := &directconnect.DescribeConnectionsInput{
		ConnectionId: &connectionId,
	}

	output, err := client.DescribeConnections(ctx, params)
	if err != nil {
		panic(err)
	}

	connectionName := output.Connections[0].ConnectionName
	fmt.Println(*connectionName) // demo-connection-001
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
