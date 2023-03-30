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
	value := "demo-internet-gateway-002"

	createInternetGatewayInput := &ec2.CreateInternetGatewayInput{
		TagSpecifications: []types.TagSpecification{
			{
				ResourceType: types.ResourceTypeInternetGateway,
				Tags: []types.Tag{
					{
						Key:   &key,
						Value: &value,
					},
				},
			},
		},
	}

	createInternetGatewayOutput, err := client.CreateInternetGateway(ctx, createInternetGatewayInput)
	if err != nil {
		panic(err)
	}

	internetGatewayId := *createInternetGatewayOutput.InternetGateway.InternetGatewayId
	fmt.Println(internetGatewayId) // igw-0f611889b41740e85

	vpcId := "vpc-019a7b633eda5caae"
	attachInternetGatewayInput := &ec2.AttachInternetGatewayInput{
		InternetGatewayId: &internetGatewayId,
		VpcId:             &vpcId,
	}

	_, err = client.AttachInternetGateway(ctx, attachInternetGatewayInput)
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
