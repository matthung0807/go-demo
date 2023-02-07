package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
)

func main() {
	ctx := context.TODO()
	client := NewDirectConnectClient(ctx)

	connectionId := "dxcon-fg5kq63s"

	virtualInterfaceName := "demo-virtual-interface-001"
	authKey := "bgp-auth-key"
	directConnectGatewayId := "e44e0dfb-82b9-4e4f-bcc1-9d196f25d0af"
	newPrivateVirtualInterface := types.NewPrivateVirtualInterface{
		Asn:                    4567,
		VirtualInterfaceName:   &virtualInterfaceName,
		Vlan:                   123,
		AddressFamily:          types.AddressFamilyIPv4,
		AuthKey:                &authKey,
		DirectConnectGatewayId: &directConnectGatewayId,
	}

	params := &directconnect.CreatePrivateVirtualInterfaceInput{
		ConnectionId:               &connectionId,
		NewPrivateVirtualInterface: &newPrivateVirtualInterface,
	}

	output, err := client.CreatePrivateVirtualInterface(ctx, params)
	if err != nil {
		panic(err)
	}

	fmt.Println(*output.VirtualInterfaceId)          // dxvif-ffwha4ij
	fmt.Println(*output.VirtualInterfaceName)        // demo-virtual-interface-001
	fmt.Println(output.VirtualInterfaceState)        // pending
	fmt.Println(*output.VirtualInterfaceType)        // private
	fmt.Println(*output.ConnectionId)                // dxcon-fg5kq63s
	fmt.Println(*output.DirectConnectGatewayId)      // e44e0dfb-82b9-4e4f-bcc1-9d196f25d0af
	fmt.Println(*output.AmazonSideAsn)               // 64512
	fmt.Println(*output.BgpPeers[0].BgpPeerId)       // dxpeer-fgzqetz9
	fmt.Println(output.BgpPeers[0].BgpPeerState)     // pending
	fmt.Println(output.BgpPeers[0].BgpStatus)        // down
	fmt.Println(*output.BgpPeers[0].AmazonAddress)   // 169.254.96.1/29
	fmt.Println(*output.BgpPeers[0].CustomerAddress) // 169.254.96.6/29
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
