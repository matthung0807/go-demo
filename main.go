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

	projectId := "project-id-1"
	region := "asia-east1"

	interconnectAttachmentService := compute.NewInterconnectAttachmentsService(service)
	vlanAttachement := &compute.InterconnectAttachment{
		Name:                   "demo-vlan-b-002",
		EdgeAvailabilityDomain: "AVAILABILITY_DOMAIN_1",
		Router:                 "https://www.googleapis.com/compute/v1/projects/project-id-1/regions/asia-east1/routers/demo-cloudrouter-002",
		Type:                   "PARTNER",
	}
	call := interconnectAttachmentService.Insert(projectId, region, vlanAttachement)
	_, err = call.Do()
	if err != nil {
		panic(err)
	}
}
