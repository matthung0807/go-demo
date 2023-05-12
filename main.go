package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3control"
)

func main() {
	ctx := context.TODO()

	client := NewS3ControlClient(ctx)

	accountId := "478900741429"
	apName := "ap-1" // access point name

	policy := `{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "AllowAllGetObject",
            "Effect": "Allow",
            "Principal": "*",
            "Action": "s3:GetObject",
            "Resource": "arn:aws:s3:ap-northeast-1:478900741429:accesspoint/ap-1/object/*"
        }
    ]
}`

	input := &s3control.PutAccessPointPolicyInput{
		AccountId: &accountId,
		Name:      &apName,
		Policy:    &policy,
	}

	_, err := client.PutAccessPointPolicy(ctx, input)
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
