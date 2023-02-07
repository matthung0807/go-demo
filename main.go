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

	connection := output.Connections[0]
	fmt.Println(*connection.ConnectionName)     // demo-connection-001
	fmt.Println(*connection.Bandwidth)          // 1Gbps
	fmt.Println(connection.ConnectionState)     // down
	fmt.Println(*connection.Location)           // CHT51
	fmt.Println(*connection.AwsLogicalDeviceId) // CHT51-2l5nybymui838
	fmt.Println(*connection.ProviderName)       // Chunghwa Telecom
	fmt.Println(connection.LoaIssueTime)        // nil
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
