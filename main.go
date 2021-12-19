package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
	accessKeyID = "AKIAIOSFODNN7EXAMPLE"                     // example
	secretKey   = "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY" // example
)

func main() {
	ctx := context.TODO()

	client := NewS3Client(ctx) // Create an Amazon S3 service client

	bucketName := "s3-demo-bucket-202112151320"
	output := GetBucketObjectOutput(ctx, client, bucketName)

	for _, object := range output.Contents {
		fmt.Printf("key=%s\n", aws.ToString(object.Key))
	}
}

func NewS3Client(ctx context.Context) *s3.Client {
	creds := aws.NewCredentialsCache(
		credentials.NewStaticCredentialsProvider(accessKeyID, secretKey, ""))

	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion("ap-northeast-1"),
		config.WithCredentialsProvider(creds),
	)
	if err != nil {
		panic(err)
	}

	return s3.NewFromConfig(cfg) // Create an Amazon S3 service client
}

func GetBucketObjectOutput(
	ctx context.Context,
	client *s3.Client,
	bucketName string) *s3.ListObjectsV2Output {

	output, err := client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		panic(err)
	}
	return output
}
