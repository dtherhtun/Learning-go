package cmd

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type ebsAction struct {
	ec2Client *ec2.Client
}

//func newClient(cfg aws.Config) *ec2Client {
//	return &ec2Client{
//		c: ec2.NewFromConfig(cfg),
//	}
//}

func (ebs *ebsAction) createSnapshot(vId string) (*string, error) {
	params := &ec2.CreateSnapshotInput{
		VolumeId: aws.String(vId),
	}

	snapshotStatus, err := ebs.ec2Client.CreateSnapshot(context.TODO(), params)
	if err != nil {
		return nil, fmt.Errorf("error occred: %w", err)
	}

	return snapshotStatus.SnapshotId, nil
}

func (ebs *ebsAction) restoreSnapshot(ssId, zone string) (*string, error) {
	params := &ec2.CreateVolumeInput{
		AvailabilityZone: aws.String(zone),
		SnapshotId:       aws.String(ssId),
		VolumeType:       "gp3",
	}

	volumeStatus, err := ebs.ec2Client.CreateVolume(context.TODO(), params)
	if err != nil {
		return nil, fmt.Errorf("error occred: %w", err)
	}

	return volumeStatus.VolumeId, nil
}

func (ebs *ebsAction) deleteVolumeSnap(id string) (string, error) {
	if strings.HasPrefix(id, "snap-") {
		params := &ec2.DeleteSnapshotInput{
			SnapshotId: aws.String(id),
		}
		_, err := ebs.ec2Client.DeleteSnapshot(context.TODO(), params)
		if err != nil {
			return "", fmt.Errorf("error occred: %w", err)
		}
		return fmt.Sprintf("Snapshot [%s] was successfully deleted.", id), nil
	}
	if strings.HasPrefix(id, "vol-") {
		params := &ec2.DeleteVolumeInput{
			VolumeId: aws.String(id),
		}
		_, err := ebs.ec2Client.DeleteVolume(context.TODO(), params)
		if err != nil {
			return "", fmt.Errorf("error occred: %w", err)
		}
		return fmt.Sprintf("Volume [%s] was successfully deleted.", id), nil
	}
	return "", errors.New("invalid input id")
}
