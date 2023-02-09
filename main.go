package main

import (
	"context"

	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	directconnecttypes "github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func NewDirectConnectClient(ctx context.Context) *directconnect.Client {
	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion("ap-northeast-1"),
	)
	if err != nil {
		panic(err)
	}

	return directconnect.NewFromConfig(cfg) // Create an Amazon Direct Connect service client
}

func NewEC2Client(ctx context.Context) *ec2.Client {
	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion("ap-northeast-1"),
	)
	if err != nil {
		panic(err)
	}

	return ec2.NewFromConfig(cfg) // Create an Amazon EC2 service client
}

func main() {
	ctx := context.TODO()
	directConnectClient := NewDirectConnectClient(ctx)
	ec2Client := NewEC2Client(ctx)

	connectionName := "demo-connection-003"
	location := "CHT51" // Chunghwa Telecom, Taipei, TWN
	bandWidth := "1Gbps"
	providerName := "Chunghwa Telecom"
	connectionId := CreateDirectConnectConnection(ctx, directConnectClient, &connectionName, &location, &bandWidth, &providerName)

	directConnectGatewayName := "demo-directconnect-gateway-003"
	amazonSideAsn := int64(64512)
	directConnectGatewayId := CreateDirectConnectGateway(ctx, directConnectClient, &directConnectGatewayName, &amazonSideAsn)

	WaitConnectionStateReady(ctx, directConnectClient, connectionId)

	virtualInterfaceName := "demo-virtual-interface-003"
	authKey := "bgp-auth-key"
	CreatePriavteVirtualInterface(ctx, directConnectClient, &virtualInterfaceName, &authKey, directConnectGatewayId, connectionId)

	virtualPrivateGatewayName := "demo-virtual-private-gateway-003"
	vpnGatewayId := CreateVirtualPrivateGateway(ctx, ec2Client, &virtualPrivateGatewayName, &amazonSideAsn)

	vpcId := "vpc-0e6e56e06a48ef314"
	AttachVirtualPrivateToVPC(ctx, ec2Client, &vpcId, vpnGatewayId)

	AssociateDirectConnectGatewayWithVirtualPrivateGateway(ctx, directConnectClient, directConnectGatewayId, vpnGatewayId)
}

func CreateDirectConnectConnection(
	ctx context.Context,
	directConnectClient *directconnect.Client,
	connectionName, location, bandWidth, providerName *string) *string {

	params := &directconnect.CreateConnectionInput{
		ConnectionName: connectionName,
		Location:       location,
		Bandwidth:      bandWidth,
		ProviderName:   providerName,
	}
	output, err := directConnectClient.CreateConnection(ctx, params)
	if err != nil {
		panic(err)
	}

	return output.ConnectionId
}

func CreateDirectConnectGateway(
	ctx context.Context,
	directConnectClient *directconnect.Client,
	directConnectGatewayName *string,
	amazonSideAsn *int64) *string {

	params := &directconnect.CreateDirectConnectGatewayInput{
		DirectConnectGatewayName: directConnectGatewayName,
		AmazonSideAsn:            amazonSideAsn,
	}
	output, err := directConnectClient.CreateDirectConnectGateway(ctx, params)
	if err != nil {
		panic(err)
	}

	return output.DirectConnectGateway.DirectConnectGatewayId
}

func WaitConnectionStateReady(ctx context.Context, directConnectClient *directconnect.Client, connectionId *string) {
	for {
		time.Sleep(time.Minute * 5) // wait 5 minute for connection state is ready
		if IsConnectionStateReady(ctx, directConnectClient, connectionId) {
			break
		}
	}
}

func IsConnectionStateReady(ctx context.Context, directConnectClient *directconnect.Client, connectionId *string) bool {
	params := &directconnect.DescribeConnectionsInput{
		ConnectionId: connectionId,
	}

	output, err := directConnectClient.DescribeConnections(ctx, params)
	if err != nil {
		panic(err)
	}

	state := output.Connections[0].ConnectionState
	return state == directconnecttypes.ConnectionStateDown ||
		state == directconnecttypes.ConnectionStateAvailable
}

func CreatePriavteVirtualInterface(
	ctx context.Context,
	directConnectClient *directconnect.Client,
	virtualInterfaceName, authKey, directConnectGatewayId, connectionId *string) {

	newPrivateVirtualInterface := directconnecttypes.NewPrivateVirtualInterface{
		Asn:                    4567,
		VirtualInterfaceName:   virtualInterfaceName,
		Vlan:                   123,
		AddressFamily:          directconnecttypes.AddressFamilyIPv4,
		AuthKey:                authKey,
		DirectConnectGatewayId: directConnectGatewayId,
	}
	params := &directconnect.CreatePrivateVirtualInterfaceInput{
		ConnectionId:               connectionId,
		NewPrivateVirtualInterface: &newPrivateVirtualInterface,
	}
	_, err := directConnectClient.CreatePrivateVirtualInterface(ctx, params)
	if err != nil {
		panic(err)
	}
}

func CreateVirtualPrivateGateway(ctx context.Context, ec2Client *ec2.Client, name *string, amazonSideAsn *int64) *string {
	key := "Name"
	tag := ec2types.Tag{
		Key:   &key,
		Value: name,
	}
	tagSpecification := ec2types.TagSpecification{
		ResourceType: ec2types.ResourceTypeVpnGateway,
		Tags:         []ec2types.Tag{tag},
	}
	params := &ec2.CreateVpnGatewayInput{
		Type:              ec2types.GatewayTypeIpsec1,
		AmazonSideAsn:     amazonSideAsn,
		TagSpecifications: []ec2types.TagSpecification{tagSpecification},
	}
	output, err := ec2Client.CreateVpnGateway(ctx, params)
	if err != nil {
		panic(err)
	}

	return output.VpnGateway.VpnGatewayId
}

func AttachVirtualPrivateToVPC(
	ctx context.Context,
	ec2Client *ec2.Client,
	vpcId, vpnGatewayId *string) {

	params := &ec2.AttachVpnGatewayInput{
		VpcId:        vpcId,
		VpnGatewayId: vpnGatewayId,
	}
	_, err := ec2Client.AttachVpnGateway(ctx, params)
	if err != nil {
		panic(err)
	}
}

func AssociateDirectConnectGatewayWithVirtualPrivateGateway(
	ctx context.Context,
	directConnectClient *directconnect.Client,
	directConnectGatewayId *string,
	vpnGatewayId *string) {

	params := &directconnect.CreateDirectConnectGatewayAssociationInput{
		DirectConnectGatewayId: directConnectGatewayId,
		VirtualGatewayId:       vpnGatewayId,
	}
	_, err := directConnectClient.CreateDirectConnectGatewayAssociation(ctx, params)
	if err != nil {
		panic(err)
	}
}
