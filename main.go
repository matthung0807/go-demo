package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	ctx := context.TODO()

	client := NewS3Client(ctx)

	bucket := "s3-demo-bucket-202112151320"
	output := GetListObjectsOutput(ctx, client, bucket)

	for _, object := range output.Contents {
		fmt.Printf("key=%s\n", aws.ToString(object.Key))
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

func GetListObjectsOutput(
	ctx context.Context,
	client *s3.Client,
	bucket string) *s3.ListObjectsV2Output {

	output, err := client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		panic(err)
	}
	return output
}
