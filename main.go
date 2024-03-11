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

	var nextToken *string
	for {
		output, err := client.GetMetricData(ctx, &cloudwatch.GetMetricDataInput{
			NextToken: nextToken,
			StartTime: aws.Time(time.Date(2024, 3, 5, 14, 0, 0, 0, time.UTC)), // 2024-03-05T14:00:00Z
			MetricDataQueries: []types.MetricDataQuery{
				{
					Id:         aws.String("q1"),
					Expression: aws.String("AVG(METRICS())"),
					Period:     aws.Int32(600),
				},
				{
					Id: aws.String("m1"),
					MetricStat: &types.MetricStat{
						Metric: &types.Metric{
							Namespace:  aws.String("AWS/EC2"),
							MetricName: aws.String("CPUUtilization"),
							Dimensions: []types.Dimension{
								{
									Name:  aws.String("InstanceId"),
									Value: aws.String("i-015766eb6a31d3413"),
								},
							},
						},
						Period: aws.Int32(600),
						Stat:   aws.String("Average"),
					},
				},
			},
		})
		if err != nil {
			panic(err)
		}
		for _, r := range output.MetricDataResults {
			if *r.Id == "q1" {
				for _, t := range r.Timestamps {
					fmt.Println(t)
				}
				for _, v := range r.Values {
					fmt.Println(v)
				}
			}
		}
		if output.NextToken == nil {
			break
		}
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
