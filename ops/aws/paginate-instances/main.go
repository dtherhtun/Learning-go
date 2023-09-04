package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
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
		MaxResults: aws.Int32(6),
	}

	paginator := ec2.NewDescribeInstancesPaginator(client, parms)

	pageCount := 1

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(context.TODO())
		if err != nil {
			fmt.Print("Error calling ec2: ", err)
		}
		fmt.Printf("Page: %v\n", pageCount)
		pageCount += 1
		for _, reservation := range page.Reservations {
			for k, instance := range reservation.Instances {
				fmt.Printf("Instance number: %v, ID: %v, Status: %v \n",
					k, *instance.InstanceId, instance.State.Name,
				)
			}
		}
	}
}
