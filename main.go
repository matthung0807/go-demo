package main

import (
	"context"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
)

func main() {
	router := gin.Default()    // get gin engine
	router.Use(cors.Default()) // allow cors
	router.GET("/cpu-utilization", HandleCpuUtilization)
	router.Run()
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

type QueryString struct {
	StartTime *time.Time `form:"startTime" binding:"required" time_format:"2006-01-02T15:04:05" time_utc:"1"`
	EndTime   *time.Time `form:"endTime" binding:"required" time_format:"2006-01-02T15:04:05" time_utc:"1"`
}

func HandleCpuUtilization(c *gin.Context) {
	ctx := c.Request.Context()
	var q QueryString
	err := c.BindQuery(&q)
	if err != nil {
		c.Error(err)
		return
	}

	client := NewCloudWatchClient(ctx)
	output, err := client.GetMetricStatistics(ctx, &cloudwatch.GetMetricStatisticsInput{
		MetricName: aws.String("CPUUtilization"),
		StartTime:  q.StartTime,
		EndTime:    q.EndTime,
		Period:     aws.Int32(300),
		Namespace:  aws.String("AWS/EC2"),
		Statistics: []types.Statistic{types.StatisticSum},
		Dimensions: []types.Dimension{
			{
				Name:  aws.String("InstanceId"),
				Value: aws.String("i-0dbd976daa9118015"),
			},
		},
	})
	if err != nil {
		c.Error(err)
		return
	}
	points := output.Datapoints
	sort.Slice(points, func(i, j int) bool {
		return points[i].Timestamp.Before(*points[j].Timestamp)
	})

	resp := Response{
		Points: lo.Map(points, func(point types.Datapoint, i int) Point {
			return Point{
				DateTime: point.Timestamp.Format("15:04"),
				Value:    strconv.FormatFloat(*point.Sum, 'f', 2, 64),
				Unit:     string(point.Unit),
			}
		}),
	}

	c.JSON(http.StatusOK, resp)
}

type Response struct {
	Points []Point `json:"points"`
}

type Point struct {
	DateTime string `json:"dateTime"`
	Value    string `json:"value"`
	Unit     string `json:"unit"`
}
