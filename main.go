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

	connectionName := "demo-connection-001"
	location := "CHT51" // Chunghwa Telecom, Taipei, TWN
	bandWidth := "1Gbps"
	providerName := "Chunghwa Telecom"
	params := &directconnect.CreateConnectionInput{
		ConnectionName: &connectionName,
		Location:       &location,
		Bandwidth:      &bandWidth,
		ProviderName:   &providerName,
	}
	output, err := client.CreateConnection(ctx, params)
	if err != nil {
		panic(err)
	}

	fmt.Println(*output.ConnectionId) // dxcon-fg5kq63s
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
