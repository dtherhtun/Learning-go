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
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/dtherhtun/Learning-go/aws/ec2revive/revive"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Alive process of ec2 and ebs",
	RunE: func(cmd *cobra.Command, args []string) error {
		stateFile := viper.GetString("state-file")
		return startProcess(stateFile)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func startProcess(stateFile string) error {
	data := &revive.Instances{}
	if err := data.Load(stateFile); err != nil {
		return err
	}

	if data.Status == "Alive" {
		return errors.New("ec2 and ebs are already alive")
	}

	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return fmt.Errorf("configuration error, %w" + err.Error())
	}
	client := revive.Client{
		EC2Client: ec2.NewFromConfig(cfg),
	}

	var instanceIDs []string

	for i := 0; i < len(data.InstanceList); i++ {
		instanceIDs = append(instanceIDs, data.InstanceList[i].ID)
		restoreResult, err := client.RestoreVolume(ctx, data.InstanceList[i].Snap, data.InstanceList[i].AZ)
		if err != nil {
			return err
		}
		if err := data.UpdateVolumeByInstanceID(data.InstanceList[i].ID, *restoreResult.VolumeId); err != nil {
			return err
		}
		fmt.Printf("Volume [%s] has been successfully restored from snapshot [%s].\n", *restoreResult.VolumeId, *restoreResult.SnapshotId)
	}

	for i := 0; i < len(data.InstanceList); i++ {
		attachResult, err := client.AttachVolume(ctx, data.InstanceList[i].ID, data.InstanceList[i].Device, data.InstanceList[i].VID)
		if err != nil {
			return err
		}
		fmt.Printf("Volume [%s] has been successfully attached to instance [%s].\n", *attachResult.VolumeId, *attachResult.InstanceId)
	}

	startResult, err := client.StartInstance(ctx, instanceIDs)
	for _, res := range startResult.StartingInstances {
		fmt.Printf("Instance [%s] is now running.\n", *res.InstanceId)
	}

	for i := 0; i < len(data.InstanceList); i++ {
		_, err := client.CleanUpSnapshot(ctx, data.InstanceList[i].Snap)
		if err != nil {
			return err
		}
		fmt.Printf("Snapshot [%s] has been successfully cleaned up.\n", data.InstanceList[i].Snap)
	}

	data.Status = "Alive"
	if err := data.Update(stateFile); err != nil {
		return err
	}
	return nil
}
