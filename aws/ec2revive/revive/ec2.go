package revive

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func (c *Client) StartInstance(ctx context.Context, instanceIDs []string) (*ec2.StartInstancesOutput, error) {

	input := &ec2.StartInstancesInput{
		InstanceIds: instanceIDs,
	}

	result, err := c.EC2Client.StartInstances(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("error starting instances: %w", err)
	}

	return result, nil
}

func (c *Client) StopInstances(ctx context.Context, instanceIDs []string) (*ec2.StopInstancesOutput, error) {

	input := &ec2.StopInstancesInput{
		InstanceIds: instanceIDs,
		Force:       aws.Bool(false),
	}

	result, err := c.EC2Client.StopInstances(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("error stopping instances: %w", err)
	}

	iw := ec2.NewInstanceStoppedWaiter(c.EC2Client)
	maxWait := 10 * time.Minute
	waitInput := ec2.DescribeInstancesInput{
		InstanceIds: instanceIDs,
	}

	if err := iw.Wait(ctx, &waitInput, maxWait); err != nil {
		return nil, fmt.Errorf("during wait for instance stop: %w", err)
	}

	return result, nil
}

func (c *Client) GetZone(ctx context.Context, instanceIDs []string) (*ec2.DescribeInstancesOutput, error) {

	input := &ec2.DescribeInstancesInput{
		InstanceIds: instanceIDs,
	}
	results, err := c.EC2Client.DescribeInstances(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("error getting instances zone: %w", err)
	}

	return results, nil
}
