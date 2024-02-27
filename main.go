package main

import (
	"context"

	compute "google.golang.org/api/compute/v1"
)

func main() {
	ctx := context.Background()
	computeService, err := compute.NewService(ctx)
	if err != nil {
		panic(err)
	}
	projectId := "project-id-1"
	zone := "asia-east2-a"

	instancesService := compute.NewInstancesService(computeService)

	instance := &compute.Instance{
		Name:               "instance-1-a",
		SourceMachineImage: "projects/project-id-1/global/machineImages/instance-1-image",
	}

	call := instancesService.Insert(projectId, zone, instance)
	_, err = call.Do()
	if err != nil {
		panic(err)
	}

}
