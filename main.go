package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

func main() {
	ctx := context.TODO()
	client := NewStsClient(ctx)

	input := &sts.GetCallerIdentityInput{}
	output, err := client.GetCallerIdentity(ctx, input)
	if err != nil {
		panic(err)
	}

	fmt.Println(*output.Account) // 123456789012
}

func NewStsClient(ctx context.Context) *sts.Client {
	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion("ap-northeast-1"),
	)
	if err != nil {
		panic(err)
	}

	return sts.NewFromConfig(cfg) // Create an Amazon STS service client
}
