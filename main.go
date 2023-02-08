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

	vpcId := "vpc-0e6e56e06a48ef314"
	vpnGatewayId := "vgw-0670c529abefaee33"
	params := &ec2.AttachVpnGatewayInput{
		VpcId:        &vpcId,
		VpnGatewayId: &vpnGatewayId,
	}
	output, err := client.AttachVpnGateway(ctx, params)
	if err != nil {
		panic(err)
	}

	fmt.Println(*output.VpcAttachment.VpcId) // vpc-0e6e56e06a48ef314
	fmt.Println(output.VpcAttachment.State)  // attaching
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
