package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3control"
)

func main() {
	ctx := context.TODO()

	client := NewS3ControlClient(ctx)

	accountId := "423456789012"
	apName := "ap-1" // access point name

	input := &s3control.GetAccessPointInput{
		AccountId: &accountId,
		Name:      &apName,
	}

	output, err := client.GetAccessPoint(ctx, input)
	if err != nil {
		panic(err)
	}

	fmt.Println(*output.AccessPointArn)    // arn:aws:s3:ap-northeast-1:423456789012:accesspoint/ap-1
	fmt.Println(*output.Alias)             // ap-1-fpebhno1smg31ehcy4heps8dkz664apn1a-s3alias
	fmt.Println(*output.BucketAccountId)   // 423456789012
	fmt.Println(*&output.NetworkOrigin)    // Internet
	fmt.Println(*&output.VpcConfiguration) // nil

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
