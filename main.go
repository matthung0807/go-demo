package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

func main() {
	ctx := context.TODO()
	client := NewLambdaClient(ctx)
	out, err := client.Invoke(ctx, &lambda.InvokeInput{
		FunctionName: aws.String("demo-func-1"),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(out.StatusCode) // 200
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
