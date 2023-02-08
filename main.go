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

	// directConnectGatewayId := "e44e0dfb-82b9-4e4f-bcc1-9d196f25d0af"
	// virtualGatewayId := "vgw-0670c529abefaee33"
	associationId := "047dd041-3388-4907-a1df-f61de4644c0b"
	params := &directconnect.DeleteDirectConnectGatewayAssociationInput{
		// DirectConnectGatewayId: &directConnectGatewayId,
		// VirtualGatewayId:       &virtualGatewayId,
		AssociationId: &associationId,
	}

	output, err := client.DeleteDirectConnectGatewayAssociation(ctx, params)
	if err != nil {
		panic(err)
	}

	fmt.Println(*output.DirectConnectGatewayAssociation.DirectConnectGatewayId) // e44e0dfb-82b9-4e4f-bcc1-9d196f25d0af
	fmt.Println(*output.DirectConnectGatewayAssociation.VirtualGatewayId)       // vgw-0670c529abefaee33
	fmt.Println(output.DirectConnectGatewayAssociation.AssociationState)        // disassociating
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
