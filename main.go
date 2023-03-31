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

	key := "Name"
	value := "demo-subnet-003"
	az := "ap-northeast-1a"

	cidrBlock := "10.1.0.0/25"
	vpcId := "vpc-019a7b633eda5caae"
	input := &ec2.CreateSubnetInput{
		VpcId:            &vpcId,
		AvailabilityZone: &az,
		CidrBlock:        &cidrBlock,
		TagSpecifications: []types.TagSpecification{
			{
				ResourceType: types.ResourceTypeSubnet,
				Tags: []types.Tag{
					{
						Key:   &key,
						Value: &value,
					},
				},
			},
		},
	}

	output, err := client.CreateSubnet(ctx, input)
	if err != nil {
		panic(err)
	}

	fmt.Println(*output.Subnet.SubnetId)                // subnet-05c7f25587be4dc58
	fmt.Println(*output.Subnet.VpcId)                   // vpc-019a7b633eda5caae
	fmt.Println(*output.Subnet.AvailableIpAddressCount) // 123
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
