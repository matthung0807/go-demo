package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func main() {
	ctx := context.TODO()

	client := NewS3Client(ctx)

	bucket := "aws-s3-bucket-202305021730" // bucket name
	output, err := client.CreateBucket(ctx, &s3.CreateBucketInput{
		Bucket: aws.String(bucket),
		CreateBucketConfiguration: &types.CreateBucketConfiguration{
			LocationConstraint: types.BucketLocationConstraintApNortheast1, // same as client config's region
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(*output.Location) // http://aws-s3-bucket-202305021730.s3.amazonaws.com/

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
