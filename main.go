package main

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/api/serviceusage/v1"
)

func main() {

	ctx := context.Background()
	serviceusageService, err := serviceusage.NewService(ctx)
	if err != nil {
		panic(err)
	}

	projectId := "project-id-1"
	name := fmt.Sprintf("projects/%s/services/serviceusage.googleapis.com", projectId)
	resp, err := serviceusageService.Services.Get(name).Do()
	if err != nil {
		panic(err)
	}
	parent := resp.Parent

	serviceNames := make([]string, 0)

	pageToken := ""
	for {
		resp, err := serviceusageService.Services.List(parent).
			Fields("nextPageToken", "services(name,state)").
			PageSize(200).
			PageToken(pageToken).
			Do()
		if err != nil {
			panic(err)
		}

		for _, s := range resp.Services {
			if strings.HasSuffix(s.Name, ".googleapis.com") {
				prefix := fmt.Sprintf("%s/services/", parent)
				serviceNames = append(serviceNames, strings.TrimPrefix(s.Name, prefix))
			}
		}

		pageToken = resp.NextPageToken
		if pageToken == "" {
			break
		}

	}

	for _, name := range serviceNames {
		fmt.Println(name)
	}

}
