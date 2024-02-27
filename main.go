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

	machineImagesService := compute.NewMachineImagesService(computeService)
	machineImage := &compute.MachineImage{
		Name:           "instance-1-image",
		SourceInstance: "projects/project-id-1/zones/asia-east2-b/instances/instance-1",
	}
	call := machineImagesService.Insert(projectId, machineImage)

	_, err = call.Do()
	if err != nil {
		panic(err)
	}

}
