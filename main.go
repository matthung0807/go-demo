package main

import (
	"context"
	"fmt"

	compute "google.golang.org/api/compute/v1"
)

func main() {
	ctx := context.Background()
	computeService, err := compute.NewService(ctx)
	if err != nil {
		panic(err)
	}
	projectId := "project-id-1"
	zone := "asia-east1-b"

	machineTypeService := compute.NewMachineTypesService(computeService)
	call := machineTypeService.List(projectId, zone)
	machineTypeList, err := call.Filter("name:e2-*").Do()
	if err != nil {
		panic(err)
	}
	for _, item := range machineTypeList.Items {
		fmt.Println(item.Name)
		fmt.Println(item.SelfLink)
	}

}
