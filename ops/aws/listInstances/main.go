package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

var client *ec2.Client

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	client = ec2.NewFromConfig(cfg)
}

func main() {
	parms := &ec2.DescribeInstancesInput{
		Filters: []types.Filter{
			{
				Name: aws.String("instance-state-name"),
				Values: []string{
					string(types.InstanceStateNameStopping),
					string(types.InstanceStateNameStopped),
					string(types.InstanceStateNameTerminated),
				},
			},
		},
		MaxResults: aws.Int32(10),
	}

	result, err := client.DescribeInstances(context.TODO(), parms)
	if err != nil {
		fmt.Println("Error calling ec2: ", err)
		return
	}

	count := 0
	for _, reservation := range result.Reservations {
		count += len(reservation.Instances)
	}

	fmt.Println("Instances: ", count)
	for _, reservation := range result.Reservations {
		for k, instance := range reservation.Instances {
			fmt.Printf("Instance number: %v, ID: %v, Status: %v \n",
				k, *instance.InstanceId, instance.State.Name,
			)
		}
	}
}
