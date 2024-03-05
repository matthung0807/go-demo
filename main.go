package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
)

func main() {
	ctx := context.TODO()
	client := NewCloudWatchClient(ctx)

	dashboardName := "cloudwatch-1"
	dashboardBody := `
{
  "widgets": []
}
`

	_, err := client.PutDashboard(ctx, &cloudwatch.PutDashboardInput{
		DashboardName: aws.String(dashboardName),
		DashboardBody: aws.String(dashboardBody),
	})
	if err != nil {
		panic(err)
	}
}

func NewCloudWatchClient(ctx context.Context) *cloudwatch.Client {
	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion("ap-northeast-1"),
	)
	if err != nil {
		panic(err)
	}

	return cloudwatch.NewFromConfig(cfg) // Create an Amazon CloudWatch service client
}
