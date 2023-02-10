package main

import (
	"context"
	"fmt"

	compute "google.golang.org/api/compute/v0.beta"
)

func main() {
	ctx := context.Background()
	service, err := compute.NewService(ctx)
	if err != nil {
		panic(err)
	}
	instancesService := compute.NewInstancesService(service)

	projectId := "project-id-1"
	zone := "asia-east1-b"
	instanceName := "demo-instance"
	call := instancesService.Get(projectId, zone, instanceName)

	instance, err := call.Do()
	if err != nil {
		panic(err)
	}

	fmt.Println(instance.Id)                    // 6192712587621936341
	fmt.Println(instance.Name)                  // demo-instance
	fmt.Println(instance.MachineType)           // https://www.googleapis.com/compute/beta/projects/project-id-1/zones/asia-east1-b/machineTypes/e2-micro
	fmt.Println(instance.Disks[0].DiskSizeGb)   // 20
	fmt.Println(instance.Disks[0].DeviceName)   // demo-instance
	fmt.Println(instance.Disks[0].Architecture) // X86_64
	fmt.Println(instance.Disks[0].Boot)         // true
}
