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
	address := &compute.Address{
		Name:         "demo-vpc-002-private-service-allocated-ip-range-003",
		IpVersion:    "IPV4",
		Region:       "GLOBAL",
		Network:      "projects/project-id-1/global/networks/demo-vpc-002",
		Address:      "10.0.0.0", // omit this field to automatic allocate range
		PrefixLength: int64(24),
		AddressType:  "INTERNAL",
		Purpose:      "VPC_PEERING",
	}
	call := globalAddressesService.Insert(projectId, address)

	_, err = call.Do()
	if err != nil {
		panic(err)
	}

}
