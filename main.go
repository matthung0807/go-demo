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
	machineImage := "instance-1-image"

	call := machineImagesService.Delete(projectId, machineImage)
	_, err = call.Do()
	if err != nil {
		panic(err)
	}

}
