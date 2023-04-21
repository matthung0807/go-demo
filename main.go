package main

import (
	"context"

	compute "google.golang.org/api/compute/v1"
)

func main() {
	ctx := context.Background()
	service, err := compute.NewService(ctx)
	if err != nil {
		panic(err)
	}

	globalAddressesService := compute.NewGlobalAddressesService(service)

	projectId := "project-id-1"
	addressResourceName := "demo-vpc-002-allocated-range-001"
	call := globalAddressesService.Delete(projectId, addressResourceName)

	_, err = call.Do()
	if err != nil {
		panic(err)
	}

}
