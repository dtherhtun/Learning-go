/*
Copyright © 2023 De Thar Htun

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/spf13/cobra"

	"github.com/dtherhtun/Learning-go/aws/ec2revive/revive"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		stopProcess()
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stopCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stopCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func stopProcess() {
	data := &revive.Instances{}
	if err := data.Load("./state.yaml"); err != nil {
		fmt.Println(err)
	}

	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	client := revive.Client{
		EC2Client: ec2.NewFromConfig(cfg),
	}

	var instanceIDs []string

	for i := 0; i < len(data.InstanceList); i++ {
		instanceIDs = append(instanceIDs, data.InstanceList[i].ID)
	}

	stopResult, err := client.StopInstances(ctx, instanceIDs)
	if err != nil {
		fmt.Println(err)
	}

	instancesInfo, err := client.GetInstancesInfo(ctx, instanceIDs)
	if err != nil {
		fmt.Println(err)
	}

	for _, res := range instancesInfo.Reservations {
		fmt.Printf("Instance [%s] has been successfully stopped.\n", *res.Instances[0].InstanceId)
		if err := data.UpdateZoneByInstanceID(*res.Instances[0].InstanceId, *res.Instances[0].Placement.AvailabilityZone); err != nil {
			fmt.Println(err)
		}
		if err := data.UpdateDeviceByInstanceID(*res.Instances[0].InstanceId, *res.Instances[0].BlockDeviceMappings[1].DeviceName); err != nil {
			fmt.Println(err)
		}
	}

	var volumeIDs []string

	for _, res := range stopResult.StoppingInstances {
		snapResult, err := client.CreateSnapshot(ctx, *res.InstanceId)
		if err != nil {
			fmt.Println(err)
		}
		if err := data.UpdateSnapshotByInstanceID(*res.InstanceId, *snapResult.Snapshots[0].SnapshotId); err != nil {
			fmt.Println(err)
		}
		volumeIDs = append(volumeIDs, *snapResult.Snapshots[0].VolumeId)
		fmt.Printf("Snapshot [%s] has completed successfully!\n", *snapResult.Snapshots[0].SnapshotId)
	}

	if err := client.RemoveVolumes(ctx, volumeIDs); err != nil {
		fmt.Println(err)
	}

	for _, v := range volumeIDs {
		fmt.Printf("Volume [%s] has been successfully deleted.\n", v)
	}

	if err := data.Update("./state.yaml"); err != nil {
		fmt.Println(err)
	}
}
