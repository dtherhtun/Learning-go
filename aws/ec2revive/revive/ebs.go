package revive

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type EC2DeleteVolumeAPI interface {
	DeleteVolume(ctx context.Context,
		params *ec2.DeleteVolumeInput,
		optFns ...func(*ec2.Options)) (*ec2.DeleteVolumeOutput, error)

	DetachVolume(ctx context.Context,
		params *ec2.DetachVolumeInput,
		optFns ...func(*ec2.Options)) (*ec2.DetachVolumeOutput, error)
}

func DeleteVolume(c context.Context, api EC2DeleteVolumeAPI, input *ec2.DeleteVolumeInput) (*ec2.DeleteVolumeOutput, error) {
	return api.DeleteVolume(c, input)
}

func DetachVolume(c context.Context, api EC2DeleteVolumeAPI, input *ec2.DetachVolumeInput) (*ec2.DetachVolumeOutput, error) {
	return api.DetachVolume(c, input)
}

func (c *Client) CreateSnapshot(ctx context.Context, instanceID string) (*ec2.CreateSnapshotsOutput, error) {

	instanceSpecification := types.InstanceSpecification{
		ExcludeBootVolume: new(bool),
		InstanceId:        aws.String(instanceID),
	}

	*instanceSpecification.ExcludeBootVolume = true

	input := &ec2.CreateSnapshotsInput{
		InstanceSpecification: &instanceSpecification,
		Description:           aws.String("Created by EC2Revive"),
		CopyTagsFromSource:    types.CopyTagsFromSourceVolume,
	}

	result, err := c.EC2Client.CreateSnapshots(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("error creating snapshot instances: %w", err)
	}

	snapshotIDs := []string{}
	if len(result.Snapshots) > 0 {
		for i := 0; i < len(result.Snapshots); i++ {
			snapshotIDs = append(snapshotIDs, *result.Snapshots[i].SnapshotId)
		}
	}

	sw := ec2.NewSnapshotCompletedWaiter(c.EC2Client)
	maxWaitTime := 5 * time.Hour
	waitInput := ec2.DescribeSnapshotsInput{
		SnapshotIds: snapshotIDs,
	}

	if err := sw.Wait(ctx, &waitInput, maxWaitTime); err != nil {
		return nil, fmt.Errorf("during wait for volumes snapshot: %w", err)
	}

	return result, nil
}

func (c *Client) RestoreVolume(ctx context.Context, snapshotID, zone string) (*ec2.CreateVolumeOutput, error) {

	input := &ec2.CreateVolumeInput{
		SnapshotId:       aws.String(snapshotID),
		AvailabilityZone: aws.String(zone),
		VolumeType:       "gp3",
	}

	result, err := c.EC2Client.CreateVolume(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("error restoring volume instances: %w", err)
	}

	aw := ec2.NewVolumeAvailableWaiter(c.EC2Client)
	maxWaitTime := 5 * time.Minute
	waitInput := ec2.DescribeVolumesInput{
		VolumeIds: []string{
			*result.VolumeId,
		},
	}
	if err := aw.Wait(ctx, &waitInput, maxWaitTime); err != nil {
		return nil, fmt.Errorf("during wait for volumes restore: %w", err)
	}

	return result, nil
}

func (c *Client) RemoveVolumes(ctx context.Context, volumeIDs []string) error {

	for _, id := range volumeIDs {
		vid := id

		inputDetach := &ec2.DetachVolumeInput{
			VolumeId: &vid,
		}

		resultDetach, err := DetachVolume(ctx, c.EC2Client, inputDetach)
		if err != nil {
			return fmt.Errorf("detach volume failed: %v,Volume Id is : %s", err, id)
		}

		aw := ec2.NewVolumeAvailableWaiter(c.EC2Client)
		maxWaitTime := 5 * time.Minute
		waitInputDetach := ec2.DescribeVolumesInput{
			VolumeIds: []string{
				*resultDetach.VolumeId,
			},
		}
		if err := aw.Wait(ctx, &waitInputDetach, maxWaitTime); err != nil {
			return fmt.Errorf("during wait for volumes detach: %w", err)
		}

		inputDelete := &ec2.DeleteVolumeInput{
			VolumeId: &vid,
		}

		_, err = DeleteVolume(ctx, c.EC2Client, inputDelete)
		if err != nil {
			return fmt.Errorf("retry delete volume action timeout:%v", err)
		}

		dw := ec2.NewVolumeDeletedWaiter(c.EC2Client)
		waitInputDelete := ec2.DescribeVolumesInput{
			VolumeIds: []string{
				*resultDetach.VolumeId,
			},
		}
		if err := dw.Wait(ctx, &waitInputDelete, maxWaitTime); err != nil {
			return fmt.Errorf("during wait for volumes delete: %w", err)
		}
	}

	return nil
}

func (c *Client) AttachVolume(ctx context.Context, instanceID, device, volumeID string) (*ec2.AttachVolumeOutput, error) {

	input := &ec2.AttachVolumeInput{
		InstanceId: aws.String(instanceID),
		Device:     aws.String(device),
		VolumeId:   aws.String(volumeID),
	}

	result, err := c.EC2Client.AttachVolume(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("error attaching volume to instance: %w", err)
	}
	iuw := ec2.NewVolumeInUseWaiter(c.EC2Client)
	maxWaitTime := 5 * time.Minute
	waitInputDetach := ec2.DescribeVolumesInput{
		VolumeIds: []string{
			*result.VolumeId,
		},
	}
	if err := iuw.Wait(ctx, &waitInputDetach, maxWaitTime); err != nil {
		return nil, fmt.Errorf("during wait for volumes inuse: %w", err)
	}

	return result, nil
}
