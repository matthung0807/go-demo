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

	// For Google services that support this functionality, this is`services/servicenetworking.googleapis.com/connections/servicenetworking-googleapis-com`.
	serviceConnectionName := "services/servicenetworking.googleapis.com/connections/servicenetworking-googleapis-com"
	vpcNetworkSelfLink := "projects/project-id-1/global/networks/demo-vpc-002" // vpc's selflink
	existedReservedPeeringRangeName := "demo-vpc-002-allocated-range-001"      // assigned allocated IP range name
	newReservedPeeringRangeName := "demo-vpc-002-allocated-range-002"          // new allocated IP range name

	connection := &servicenetworking.Connection{
		Network: vpcNetworkSelfLink,
		ReservedPeeringRanges: []string{
			existedReservedPeeringRangeName,
			newReservedPeeringRangeName,
		},
	}

	call := serviceConnectionService.Patch(serviceConnectionName, connection)
	_, err = call.Do()
	if err != nil {
		panic(err)
	}

}
