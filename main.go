package main

import (
	"context"

	"google.golang.org/api/servicenetworking/v1"
)

func main() {
	ctx := context.Background()
	service, err := servicenetworking.NewService(ctx)
	if err != nil {
		panic(err)
	}

	serviceConnectionService := servicenetworking.NewServicesConnectionsService(service)

	parent := "services/servicenetworking.googleapis.com"                      // For Google services that support this functionality, this value is `services/servicenetworking.googleapis.com`.
	vpcNetworkSelfLink := "projects/project-id-1/global/networks/demo-vpc-002" // vpc's selflink
	reservedPeeringRangeName := "demo-vpc-002-allocated-range-001"             // allocated IP range name

	connection := &servicenetworking.Connection{
		Network:               vpcNetworkSelfLink,
		ReservedPeeringRanges: []string{reservedPeeringRangeName},
	}

	call := serviceConnectionService.Create(parent, connection)
	_, err = call.Do()
	if err != nil {
		panic(err)
	}

}
