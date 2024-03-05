package main

import (
	"context"
	"fmt"

	compute "google.golang.org/api/compute/v1"
)

func main() {
	ctx := context.Background()
	computeService, err := compute.NewService(ctx)
	if err != nil {
		panic(err)
	}
	projectId := "project-id-1"

	machineImagesService := compute.NewMachineImagesService(computeService)
	machineImage := "instance-1-image"
	call := machineImagesService.Get(projectId, machineImage)

	result, err := call.Do()
	if err != nil {
		panic(err)
	}
	fmt.Println(result.SelfLink)
	fmt.Println(result.SourceInstance)
	fmt.Println(result.TotalStorageBytes)

	for i, saveDisk := range result.SavedDisks {
		fmt.Printf("=====saveDisk[%d]=====\n", i)
		fmt.Println(saveDisk.SourceDisk)
		fmt.Println(saveDisk.StorageBytes)
	}

	for i, disk := range result.InstanceProperties.Disks {
		fmt.Printf("=====disk[%d]=====\n", i)
		fmt.Println(disk.Index)
		fmt.Println(disk.Boot)
		fmt.Println(disk.DeviceName)
		fmt.Println(disk.DiskSizeGb)
	}

}
