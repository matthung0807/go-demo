package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func main() {
	ctx := context.TODO()
	client := NewEC2Client(ctx)

	instanceName := "demo-ec2-instance-002"
	imageId := "ami-0e0820ad173f20fbb" // Amazon Linux 2023 AMI 2023.0.20230503.0 x86_64 HVM kernel-6.1
	KeyName := "demo-key"              // key pairs name
	securityGroupId := "sg-0e44d58befeb89923"
	subnetId := "subnet-05c7f25587be4dc58"

	input := &ec2.RunInstancesInput{
		MaxCount:       aws.Int32(1),
		MinCount:       aws.Int32(1),
		DisableApiStop: aws.Bool(false),
		ImageId:        &imageId,
		InstanceType:   types.InstanceTypeT2Micro, // t2.micro
		KeyName:        &KeyName,
		NetworkInterfaces: []types.InstanceNetworkInterfaceSpecification{
			{
				AssociatePublicIpAddress: aws.Bool(true),
				DeleteOnTermination:      aws.Bool(true),
				DeviceIndex:              aws.Int32(0),
				Groups: []string{
					securityGroupId,
				},
				InterfaceType:    aws.String("interface"),
				NetworkCardIndex: aws.Int32(0),
				SubnetId:         &subnetId,
			},
		},
		TagSpecifications: []types.TagSpecification{
			{
				ResourceType: types.ResourceTypeInstance,
				Tags: []types.Tag{
					{
						Key:   aws.String("Name"),
						Value: aws.String(instanceName),
					},
				},
			},
		},
	}
	output, err := client.RunInstances(ctx, input)
	if err != nil {
		panic(err)
	}

	waiter := ec2.NewInstanceRunningWaiter(client)
	describeInput := &ec2.DescribeInstancesInput{
		InstanceIds: []string{aws.ToString(output.Instances[0].InstanceId)},
	}
	maxWaitDur := 60 * time.Second
	describeOutput, err := waiter.WaitForOutput(ctx, describeInput, maxWaitDur)
	if err != nil {
		panic(err)
	}

	instance := describeOutput.Reservations[0].Instances[0]
	fmt.Println(*instance.InstanceId)                              // i-0dbcc73ace6b52d42
	fmt.Println(*instance.SubnetId)                                // subnet-05c7f25587be4dc58
	fmt.Println(*instance.NetworkInterfaces[0].NetworkInterfaceId) // eni-0cd6cc75c832083f3
	fmt.Println(*instance.PublicIpAddress)                         // 3.112.238.189
	fmt.Println(*instance.PrivateIpAddress)                        // 10.1.0.53
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
