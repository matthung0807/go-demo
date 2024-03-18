package main

import (
	"context"
	"fmt"
	"sort"
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
		MetricName: aws.String("VolumeReadBytes"),
		StartTime:  aws.Time(time.Now().Add(-time.Hour * 1)),
		EndTime:    aws.Time(time.Now()),
		Period:     aws.Int32(600),
		Namespace:  aws.String("AWS/EBS"),
		Statistics: []types.Statistic{types.StatisticSum},
		Dimensions: []types.Dimension{
			{
				Name:  aws.String("VolumeId"),
				Value: aws.String("vol-02a899c6436c371a5"),
			},
		},
	})
	if err != nil {
		panic(err)
	}

	dataPoints := output.Datapoints
	sort.Slice(dataPoints, func(i, j int) bool {
		return dataPoints[i].Timestamp.Before(*dataPoints[j].Timestamp)
	})

	for _, point := range output.Datapoints {
		fmt.Println(*point.Timestamp)
		fmt.Println(*point.Sum)
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
