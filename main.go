package main

import (
	"context"
	"fmt"
	"strings"

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
	call := globalAddressesService.List(projectId)

	addressList, err := call.Do()
	if err != nil {
		panic(err)
	}

	vpcSelfLink := "projects/proejct-id-1/global/networks/demo-vpc-002"
	for _, item := range addressList.Items {
		if strings.Contains(item.Network, vpcSelfLink) {
			fmt.Println(item.Network)
			fmt.Println(item.Name)
			fmt.Printf("%s/%d\n", item.Address, item.PrefixLength)
		}
	}

}
