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
	routersService := compute.NewRoutersService(service)

	projectId := "project-id-1"
	region := "asia-east1"
	router := &compute.Router{
		Name:    "demo-cloudrouter-002",
		Network: "projects/project-id-1/global/networks/demo-vpc-001",
		Bgp: &compute.RouterBgp{
			AdvertiseMode: "DEFAULT",
			Asn:           int64(16550),
		},
	}

	call := routersService.Insert(projectId, region, router)
	_, err = call.Do()
	if err != nil {
		panic(err)
	}
}
