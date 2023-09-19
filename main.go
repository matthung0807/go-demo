package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
)

func main() {
	ctx := context.TODO()
	client := NewLambdaClient(ctx)
	out, err := client.CreateFunction(ctx, &lambda.CreateFunctionInput{
		Code: &types.FunctionCode{
			S3Bucket: aws.String("s3-demo-bucket-202309191740"),
			S3Key:    aws.String("demo-func-1.zip"),
		},
		FunctionName: aws.String("demo-func-1"),
		Role:         aws.String("arn:aws:iam::478900741429:role/lambda-basic-execution-role-1"),
		Architectures: []types.Architecture{
			types.ArchitectureX8664,
		},
		Handler:     aws.String("bootstrap"),
		Runtime:     types.RuntimeGo1x,
		PackageType: types.PackageTypeZip,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(*out.FunctionArn) // 200
}

func NewLambdaClient(ctx context.Context) *lambda.Client {
	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion("ap-northeast-1"),
	)
	if err != nil {
		panic(err)
	}

	return lambda.NewFromConfig(cfg) // Create an Amazon Lambda service client
}
