package main

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	ctx := context.TODO()
	client := NewS3Client(ctx)
	input := CreateInput()

	_, err := PutFile(ctx, client, input)
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

func CreateInput() *s3.PutObjectInput {
	bucket := "s3-demo-bucket-202112151320"
	filename := "greeting.txt"
	reader := strings.NewReader("good day")

	return &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &filename,
		Body:   reader,
	}
}

func PutFile(
	ctx context.Context,
	client *s3.Client,
	input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {

	return client.PutObject(ctx, input)
}
