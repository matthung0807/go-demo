package main

import (
	"context"
	"fmt"

	"google.golang.org/api/serviceusage/v1"
)

func main() {
	ctx := context.Background()
	serviceusageService, err := serviceusage.NewService(ctx)
	if err != nil {
		panic(err)
	}

	projectId := "project-id-1"
	serviceName := "apigateway.googleapis.com" // API Gateway API service name
	name := fmt.Sprintf("projects/%s/services/%s", projectId, serviceName)
	req := &serviceusage.EnableServiceRequest{}

	_, err = serviceusageService.Services.Enable(name, req).Do()
	if err != nil {
		panic(err)
	}

}
