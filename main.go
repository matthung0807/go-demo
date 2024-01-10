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
	projectId := "debian-cloud"

	imagesService := compute.NewImagesService(computeService)

	call := imagesService.List(projectId)
	imageList, err := call.
		Filter("(name:debian-12-*) AND (family:debian-12)").
		Do()
	if err != nil {
		panic(err)
	}

	for _, item := range imageList.Items {
		fmt.Println(item.Name)
		fmt.Println(item.SelfLink)
	}

}
