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

	directConnectGatewayName := "demo-directconnect-gateway-001"
	amazonSideAsn := int64(64512)
	params := &directconnect.CreateDirectConnectGatewayInput{
		DirectConnectGatewayName: &directConnectGatewayName,
		AmazonSideAsn:            &amazonSideAsn,
	}

	output, err := client.CreateDirectConnectGateway(ctx, params)
	if err != nil {
		panic(err)
	}

	dcg := output.DirectConnectGateway

	fmt.Println(*dcg.DirectConnectGatewayId)   // e44e0dfb-82b9-4e4f-bcc1-9d196f25d0af
	fmt.Println(dcg.DirectConnectGatewayState) // available
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
