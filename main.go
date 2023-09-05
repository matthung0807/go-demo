package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/pricing"
)

func main() {
	ctx := context.TODO()
	client := NewPricingClient(ctx)
	out, err := client.DescribeServices(ctx, &pricing.DescribeServicesInput{})
	if err != nil {
		panic(err)
	}

	for _, s := range out.Services {
		fmt.Println(*s.ServiceCode)
	}
}

func NewPricingClient(ctx context.Context) *pricing.Client {
	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion("ap-south-1"),
	)
	if err != nil {
		panic(err)
	}

	return pricing.NewFromConfig(cfg) // Create an Pricing service client
}
