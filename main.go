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
			ImageUri: aws.String("123456789012.dkr.ecr.ap-northeast-1.amazonaws.com/ecr-1:latest"),
		},
		FunctionName: aws.String("demo-func-1"),
		Role:         aws.String("arn:aws:iam::123456789012:role/lambda-basic-execution-role-1"),
		Architectures: []types.Architecture{
			types.ArchitectureX8664,
		},
		PackageType: types.PackageTypeImage,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(*out.FunctionArn) // arn:aws:lambda:ap-northeast-1:123456789012:function:demo-func-1
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
