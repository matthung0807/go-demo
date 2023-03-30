package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func main() {
	ctx := context.TODO()
	client := NewEC2Client(ctx)

	vpcId := "vpc-019a7b633eda5caae"
	key := "Name"
	value := "demo-route-table-002"

	input := &ec2.CreateRouteTableInput{
		VpcId: &vpcId,
		TagSpecifications: []types.TagSpecification{
			{
				ResourceType: types.ResourceTypeRouteTable,
				Tags: []types.Tag{
					{
						Key:   &key,
						Value: &value,
					},
				},
			},
		},
	}

	output, err := client.CreateRouteTable(ctx, input)

	if err != nil {
		panic(err)
	}
	fmt.Println(*output.RouteTable.VpcId)        // vpc-019a7b633eda5caae
	fmt.Println(*output.RouteTable.RouteTableId) // rtb-0b0e21c8e3b1cda13
	if len(output.RouteTable.Routes) > 0 {
		fmt.Println(*output.RouteTable.Routes[0].DestinationCidrBlock) // 10.1.0.0/24
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
