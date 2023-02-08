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

	amazonSideAsn := int64(64512)
	key := "Name"
	value := "demo-virtual-private-gateway-001"
	tag := types.Tag{
		Key:   &key,
		Value: &value,
	}
	tagSpecification := types.TagSpecification{
		ResourceType: types.ResourceTypeVpnGateway,
		Tags:         []types.Tag{tag},
	}
	params := &ec2.CreateVpnGatewayInput{
		Type:              types.GatewayTypeIpsec1,
		AmazonSideAsn:     &amazonSideAsn,
		TagSpecifications: []types.TagSpecification{tagSpecification},
	}

	output, err := client.CreateVpnGateway(ctx, params)
	if err != nil {
		panic(err)
	}

	vpg := output.VpnGateway
	fmt.Println(*vpg.VpnGatewayId) // vgw-0670c529abefaee33
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
