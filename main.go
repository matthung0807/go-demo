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

	routeTableId := "rtb-0b0e21c8e3b1cda13"
	subnetId := "subnet-05c7f25587be4dc58"

	input := &ec2.AssociateRouteTableInput{
		RouteTableId: &routeTableId,
		SubnetId:     &subnetId,
	}

	output, err := client.AssociateRouteTable(ctx, input)
	if err != nil {
		panic(err)
	}

	fmt.Println(*output.AssociationId) // rtbassoc-0d2dbf31192e46e91
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
