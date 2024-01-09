package main

import (
	"context"

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

	instancesService := compute.NewInstancesService(computeService)

	attachDisks := []*compute.AttachedDisk{
		{
			AutoDelete: true,
			Boot:       true,
			InitializeParams: &compute.AttachedDiskInitializeParams{
				DiskName:    "instance-1-boot-disk",
				DiskSizeGb:  20,
				DiskType:    "projects/project-id-1/zones/asia-east1-b/diskTypes/pd-balanced",
				SourceImage: "projects/debian-cloud/global/images/debian-11-bullseye-arm64-v20231212", // OS Image
			},
		},
	}

	instance := &compute.Instance{
		Name:               "instance-1",
		DeletionProtection: false,
		MachineType:        "https://www.googleapis.com/compute/v1/projects/project-id-1/zones/asia-east1-b/machineTypes/e2-small",
		NetworkInterfaces: []*compute.NetworkInterface{
			{
				StackType:  "IPV4_ONLY",
				Subnetwork: "projects/project-id-1/regions/asia-east1/subnetworks/test-east1-b",
			},
		},
		Disks: attachDisks,
	}

	call := instancesService.Insert(projectId, zone, instance)
	_, err = call.Do()
	if err != nil {
		panic(err)
	}

}
