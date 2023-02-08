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
	virtualGatewayId := "vgw-0670c529abefaee33"
	params := &directconnect.DescribeDirectConnectGatewayAssociationsInput{
		DirectConnectGatewayId: &directConnectGatewayId,
		VirtualGatewayId:       &virtualGatewayId,
	}

	output, err := client.DescribeDirectConnectGatewayAssociations(ctx, params)
	if err != nil {
		panic(err)
	}

	fmt.Println(*output.DirectConnectGatewayAssociations[0].AssociationId)   // 047dd041-3388-4907-a1df-f61de4644c0b
	fmt.Println(output.DirectConnectGatewayAssociations[0].AssociationState) // associated
	for _, prefix := range output.DirectConnectGatewayAssociations[0].AllowedPrefixesToDirectConnectGateway {
		fmt.Println(*prefix.Cidr)
	}
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
