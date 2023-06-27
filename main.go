package main

import (
	"context"
	"fmt"

	"google.golang.org/api/cloudresourcemanager/v1"
)

func main() {
	ctx := context.Background()
	crmService, err := cloudresourcemanager.NewService(ctx)
	if err != nil {
		panic(err)
	}

	projectId := "allen-test-312507"
	project, err := crmService.Projects.Get(projectId).Do()
	if err != nil {
		panic(err)
	}
	fmt.Println(project.Name)          // gcp-demo
	fmt.Println(project.ProjectId)     // project-id-1
	fmt.Println(project.ProjectNumber) // 948024786833
}
