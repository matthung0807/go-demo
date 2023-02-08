package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func main() {
	ctx := context.TODO()
	client := NewEC2Client(ctx)

	vpnGatewayId := "vgw-0670c529abefaee33"
	params := &ec2.DescribeVpnGatewaysInput{
		VpnGatewayIds: []string{vpnGatewayId},
	}

	output, err := client.DescribeVpnGateways(ctx, params)
	if err != nil {
		panic(err)
	}

	fmt.Println(*output.VpnGateways[0].VpcAttachments[0].VpcId) // vpc-0e6e56e06a48ef314
}

func NewEC2Client(ctx context.Context) *ec2.Client {
	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion("ap-northeast-1"),
	)
	if err != nil {
		panic(err)
	}

	return ec2.NewFromConfig(cfg) // Create an Amazon EC2 service client
}
