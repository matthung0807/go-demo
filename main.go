package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	ctx := context.TODO()

	client := NewS3Client(ctx)

	bucket := "s3-demo-bucket-202112151320"
	key := "hello.txt"
	in := CreateGetObjectInput(&bucket, &key)
	content := GetObjectContent(ctx, client, in)
	fmt.Printf("content=%s\n", content)
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

func CreateGetObjectInput(bucket *string, key *string) *s3.GetObjectInput {
	return &s3.GetObjectInput{
		Bucket: bucket,
		Key:    key,
	}
}

func GetObjectContent(ctx context.Context, client *s3.Client, input *s3.GetObjectInput) string {
	output, err := client.GetObject(ctx, input)
	if err != nil {
		panic(err)
	}

	b, err := ioutil.ReadAll(output.Body)
	if err != nil {
		panic(err)
	}
	defer output.Body.Close()

	return string(b)
}
