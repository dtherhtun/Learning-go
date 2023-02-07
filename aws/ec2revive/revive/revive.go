package revive

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"gopkg.in/yaml.v2"
)

type Client struct {
	EC2Client *ec2.Client
}

// Instance struct representing an EC2 instance
type Instance struct {
	ID     string `yaml:"id"`
	AZ     string `yaml:"az"`
	VID    string `yaml:"vid"`
	Snap   string `yaml:"snap"`
	Device string `yaml:"device"`
}

// Instances struct representing a list of EC2 instances
type Instances struct {
	InstanceList []Instance `yaml:"instances"`
}

func (il *Instances) Load(statefile string) error {
	f, err := os.Open(statefile)
	if err != nil {
		return fmt.Errorf("error reading YAML file: %v\n", err)
	}
	defer f.Close()
	err = yaml.NewDecoder(f).Decode(il)
	if err != nil {

		return fmt.Errorf("error unmarshaling YAML: %v\n", err)
	}
	return nil
}

func (il *Instances) Update(statefile string) error {
	f, err := os.Create(statefile)
	if err != nil {
		return fmt.Errorf("error reading YAML file: %v\n", err)
	}
	defer f.Close()
	err = yaml.NewEncoder(f).Encode(il)
	if err != nil {

		return fmt.Errorf("error marshaling YAML: %v\n", err)
	}
	return nil
}

func (il *Instances) UpdateSnapshotByInstanceID(instanceID, snapshotID string) error {
	for i := range il.InstanceList {
		if il.InstanceList[i].ID == instanceID {
			il.InstanceList[i].Snap = snapshotID
			return nil
		}
	}
	return fmt.Errorf("instance with ID %s not found", instanceID)
}

func (il *Instances) UpdateZoneByInstanceID(instanceID, zone string) error {
	for i := range il.InstanceList {
		if il.InstanceList[i].ID == instanceID {
			il.InstanceList[i].AZ = zone
			return nil
		}
	}
	return fmt.Errorf("instance with ID %s not found", instanceID)
}

func (il *Instances) UpdateVolumeByInstanceID(instanceID, volumeID string) error {
	for i := range il.InstanceList {
		if il.InstanceList[i].ID == instanceID {
			il.InstanceList[i].VID = volumeID
			return nil
		}
	}
	return fmt.Errorf("instance with ID %s not found", instanceID)
}

func (il *Instances) UpdateDeviceByInstanceID(instanceID, device string) error {
	for i := range il.InstanceList {
		if il.InstanceList[i].ID == instanceID {
			il.InstanceList[i].Device = device
			return nil
		}
	}
	return fmt.Errorf("instance with ID %s not found", instanceID)
}
