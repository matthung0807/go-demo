package main

import (
	"context"
	"fmt"
	"time"

	monitoring "cloud.google.com/go/monitoring/apiv3/v2"
	"cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/types/known/durationpb"
)

func main() {
	ctx := context.Background()
	c, err := monitoring.NewMetricClient(ctx)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	projectId := "project-id-1"

	now := time.Now().UTC()
	startTime := now.Add(time.Minute * -10).Unix()
	endTime := now.Unix()

	req := &monitoringpb.ListTimeSeriesRequest{
		Name: "projects/" + projectId,
		Filter: `
	resource.type="gce_instance"
	metric.type="compute.googleapis.com/instance/cpu/utilization" AND
	metric.labels.instance_name = "instance-1"
	`,
		Interval: &monitoringpb.TimeInterval{
			StartTime: &timestamp.Timestamp{Seconds: startTime},
			EndTime:   &timestamp.Timestamp{Seconds: endTime},
		},
		Aggregation: &monitoringpb.Aggregation{
			AlignmentPeriod:  durationpb.New(time.Minute * 10),
			PerSeriesAligner: monitoringpb.Aggregation_ALIGN_SUM,
		},
	}
	it := c.ListTimeSeries(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			panic(err)
		}

		for _, p := range resp.Points {
			fmt.Println(p.GetInterval().GetEndTime().AsTime())
			fmt.Println(p.GetValue().GetDoubleValue())
		}
	}

}
