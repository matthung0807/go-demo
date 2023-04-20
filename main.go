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

	networksService := compute.NewNetworksService(service)

	projectId := "project-id-1"
	vpcName := "demo-vpc-002"
	network := &compute.Network{
		AutoCreateSubnetworks: true,
		Mtu:                   int64(1460),
		Name:                  vpcName,
	}
	call := networksService.Insert(projectId, network)

	op, err := call.Do()
	if err != nil {
		panic(err)
	}

	_, err = service.GlobalOperations.
		Wait(projectId, op.Name).
		Do()
	if err != nil {
		panic(err)
	}
}
