package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func main() {
	ctx := context.TODO()
	client := NewEC2Client(ctx)

	routeTableId := "rtb-0b0e21c8e3b1cda13"
	cidrBlock := "0.0.0.0/0"
	internetGatewayId := "igw-0f611889b41740e85"

	input := &ec2.CreateRouteInput{
		RouteTableId:         &routeTableId,
		DestinationCidrBlock: &cidrBlock,
		GatewayId:            &internetGatewayId,
	}

	_, err := client.CreateRoute(ctx, input)

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
