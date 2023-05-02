package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	ctx := context.TODO()

	client := NewS3Client(ctx)

	policy := `{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "DenyAllGetObject",
            "Effect": "Deny",
            "Principal": {
                "AWS": "*"
            },
            "Action": "s3:GetObject",
            "Resource": "arn:aws:s3:::aws-s3-bucket-202305021730/*"
        }
    ]
}
`
	bucket := "aws-s3-bucket-202305021730" // bucket name
	input := &s3.PutBucketPolicyInput{
		Bucket: &bucket,
		Policy: &policy,
	}
	_, err := client.PutBucketPolicy(ctx, input)
	if err != nil {
		panic(err)
	}
}

func NewS3Client(ctx context.Context) *s3.Client {
	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion("ap-northeast-1"),
	)
	if err != nil {
		panic(err)
	}
	return s3.NewFromConfig(cfg) // Create an Amazon S3 service client
}
