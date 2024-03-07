package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
)

func main() {
	ctx := context.TODO()
	client := NewCloudWatchClient(ctx)

	output, err := client.GetMetricStatistics(ctx, &cloudwatch.GetMetricStatisticsInput{
		MetricName: aws.String("CPUUtilization"),
		StartTime:  aws.Time(time.Date(2024, 3, 5, 14, 0, 0, 0, time.UTC)), // 2024-03-05T14:00:00Z
		EndTime:    aws.Time(time.Date(2024, 3, 5, 15, 0, 0, 0, time.UTC)), // 2024-03-05T15:00:00Z
		Period:     aws.Int32(600),
		Namespace:  aws.String("AWS/EC2"),
		Statistics: []types.Statistic{types.StatisticAverage},
		Dimensions: []types.Dimension{
			{
				Name:  aws.String("InstanceId"),
				Value: aws.String("i-015766eb6a31d3413"),
			},
		},
	})
	if err != nil {
		panic(err)
	}

	for _, point := range output.Datapoints {
		fmt.Println(*point.Timestamp)
		fmt.Println(*point.Average)
		fmt.Println(point.Unit)
		fmt.Println("========================================")
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
