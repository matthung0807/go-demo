package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/pricing"
	"github.com/aws/aws-sdk-go-v2/service/pricing/types"
)

func main() {
	ctx := context.TODO()
	client := NewPricingClient(ctx)

	serviceCode := "AmazonEC2"
	out, err := client.GetProducts(ctx, &pricing.GetProductsInput{
		ServiceCode: &serviceCode,
		MaxResults:  aws.Int32(2),
		Filters: []types.Filter{
			{
				Field: aws.String("sku"),
				Type:  "TERM_MATCH",
				Value: aws.String("7UV2VKCDN6WVVP54"),
			},
		},
	})
	if err != nil {
		panic(err)
	}

	for _, price := range out.PriceList {
		fmt.Println(price)
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
