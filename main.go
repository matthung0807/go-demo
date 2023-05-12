package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3control"
)

func main() {
	ctx := context.TODO()

	client := NewS3ControlClient(ctx)

	accountId := "423456789012"
	apName := "ap-1" // access point name

	input := &s3control.DeleteAccessPointPolicyInput{
		AccountId: &accountId,
		Name:      &apName,
	}

	_, err := client.DeleteAccessPointPolicy(ctx, input)
	if err != nil {
		panic(err)
	}

}

func NewS3ControlClient(ctx context.Context) *s3control.Client {
	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion("ap-northeast-1"),
	)
	if err != nil {
		panic(err)
	}
	return s3control.NewFromConfig(cfg) // Create an Amazon S3 Control client
}
