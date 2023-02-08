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
	params := &directconnect.DescribeVirtualInterfacesInput{
		ConnectionId: &connectionId,
	}

	output, err := client.DescribeVirtualInterfaces(ctx, params)
	if err != nil {
		panic(err)
	}

	fmt.Println(*output.VirtualInterfaces[0].VirtualInterfaceId)          // dxvif-ffwha4ij
	fmt.Println(*output.VirtualInterfaces[0].VirtualInterfaceName)        // demo-virtual-interface-001
	fmt.Println(output.VirtualInterfaces[0].VirtualInterfaceState)        // down
	fmt.Println(*output.VirtualInterfaces[0].VirtualInterfaceType)        // private
	fmt.Println(*output.VirtualInterfaces[0].DirectConnectGatewayId)      // e44e0dfb-82b9-4e4f-bcc1-9d196f25d0af
	fmt.Println(*output.VirtualInterfaces[0].AmazonSideAsn)               // 64512
	fmt.Println(*output.VirtualInterfaces[0].BgpPeers[0].BgpPeerId)       // dxpeer-fgzqetz9
	fmt.Println(output.VirtualInterfaces[0].BgpPeers[0].BgpPeerState)     // available
	fmt.Println(output.VirtualInterfaces[0].BgpPeers[0].BgpStatus)        // down
	fmt.Println(*output.VirtualInterfaces[0].BgpPeers[0].AmazonAddress)   // 169.254.96.1/29
	fmt.Println(*output.VirtualInterfaces[0].BgpPeers[0].CustomerAddress) // 169.254.96.6/29
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
