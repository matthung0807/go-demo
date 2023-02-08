package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func main() {
	ctx := context.TODO()
	client := NewEC2Client(ctx)

	vpnGatewayId := "vgw-0670c529abefaee33"
	vpnId := "vpc-0e6e56e06a48ef314"
	params := &ec2.DetachVpnGatewayInput{
		VpnGatewayId: &vpnGatewayId,
		VpcId:        &vpnId,
	}

	_, err := client.DetachVpnGateway(ctx, params)
	if err != nil {
		panic(err)
	}
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
