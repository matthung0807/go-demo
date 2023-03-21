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

	routerName := "demo-cloudrouter-002"
	router := &compute.Router{
		Name:    routerName,
		Network: "projects/project-id-1/global/networks/demo-vpc-001",
		Bgp: &compute.RouterBgp{
			AdvertiseMode: "DEFAULT",
			Asn:           int64(16550),
		},
	}

	op, err := routersService.Insert(projectId, region, router).Do()
	if err != nil {
		panic(err)
	}

	_, err = service.RegionOperations.Wait(projectId, region, op.Name).Do()
	if err != nil {
		panic(err)
	}

	router, err = service.Routers.Get(projectId, region, routerName).Do()
	if err != nil {
		panic(err)
	}

	interconnectAttachmentService := compute.NewInterconnectAttachmentsService(service)
	vlanAttachement := &compute.InterconnectAttachment{
		Name:                   "demo-vlan-b-002",
		EdgeAvailabilityDomain: "AVAILABILITY_DOMAIN_1",
		Router:                 router.SelfLink,
		Type:                   "PARTNER",
	}
	op, err = interconnectAttachmentService.Insert(projectId, region, vlanAttachement).Do()
	if err != nil {
		panic(err)
	}
	_, err = service.RegionOperations.Wait(projectId, region, op.Name).Do()
	if err != nil {
		panic(err)
	}
}
