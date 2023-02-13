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
	routersService := compute.NewRoutersService(service)

	projectId := "allen-test-312507"
	region := "asia-east1"
	routerName := "demo-cloudrouter-002"
	call := routersService.Get(projectId, region, routerName)
	router, err := call.Do()
	if err != nil {
		panic(err)
	}

	fmt.Println(router.Name)              // demo-cloudrouter-002
	fmt.Println(router.Kind)              // compute#router
	fmt.Println(router.Bgp.Asn)           // 16550
	fmt.Println(router.Bgp.AdvertiseMode) // DEFAULT
	fmt.Println(router.Network)           // https://www.googleapis.com/compute/v1/projects/project-id-1/global/networks/demo-vpc-001
	fmt.Println(router.SelfLink)          // https://www.googleapis.com/compute/v1/projects/project-id-1/regions/asia-east1/routers/demo-cloudrouter-002
}
