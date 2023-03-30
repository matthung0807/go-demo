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
	value := "demo-vpc-002"

	cidrBlock := "10.1.0.0/24"
	input := &ec2.CreateVpcInput{
		CidrBlock: &cidrBlock,
		TagSpecifications: []types.TagSpecification{
			{
				ResourceType: types.ResourceTypeVpc,
				Tags: []types.Tag{
					{
						Key:   &key,
						Value: &value,
					},
				},
			},
		},
	}
	output, err := client.CreateVpc(ctx, input)
	if err != nil {
		panic(err)
	}
	fmt.Println(*output.Vpc.VpcId)     // vpc-019a7b633eda5caae
	fmt.Println(*output.Vpc.CidrBlock) // 10.1.0.0/24
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
