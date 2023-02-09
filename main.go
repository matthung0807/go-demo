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

	directConnectGatewayId := "e44e0dfb-82b9-4e4f-bcc1-9d196f25d0af"
	params := &directconnect.DeleteDirectConnectGatewayInput{
		DirectConnectGatewayId: &directConnectGatewayId,
	}

	output, err := client.DeleteDirectConnectGateway(ctx, params)
	if err != nil {
		panic(err)
	}

	dcg := output.DirectConnectGateway
	fmt.Println(dcg.DirectConnectGatewayState) // deleting
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
