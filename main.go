package main

import (
	"context"
	"fmt"

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
	call := networksService.Get(projectId, vpcName)

	network, err := call.Do()
	if err != nil {
		panic(err)
	}

	fmt.Println(network.Name)             // demo-vpc-002
	fmt.Println(len(network.Subnetworks)) // 37
	fmt.Println(network.Subnetworks[0])   // https://www.googleapis.com/compute/v1/projects/project-id-1/regions/europe-west9/subnetworks/demo-vpc-002
}
